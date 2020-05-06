package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/xerrors"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Auth struct {
	domain              string
	audience            string
	debug               bool
	credentialsOptional bool
}

func New(domain, audience string, debug, credentialsOptional bool) *Auth {
	return &Auth{domain, audience, debug, credentialsOptional}
}

func (a *Auth) CheckJWTHandlerFunc() func(next http.Handler) http.Handler {
	middleware := jwtMiddleware.New(jwtMiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(a.audience, false)
			if !checkAud {
				return token, xerrors.New("Invalid audience.")
			}
			// Verify 'iss' claim
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(fmt.Sprintf("https://%s/", a.domain), false)
			if !checkIss {
				return token, xerrors.New("Invalid issuer.")
			}

			cert, err := a.getPemCert(token)
			if err != nil {
				return token, xerrors.New("Invalid token.")
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil
		},
		SigningMethod:       jwt.SigningMethodRS256,
		Debug:               a.debug,
		CredentialsOptional: a.credentialsOptional,
	})
	return middleware.Handler
}

func (a *Auth) getPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get(fmt.Sprintf("https://%s/.well-known/jwks.json", a.domain))

	if err != nil {
		return cert, xerrors.Errorf("failed to http.Get: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, xerrors.Errorf("failed to decode jwks: %w", err)
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, xerrors.New("cert is blank.")
	}

	return cert, nil
}

type CustomClaims struct {
	Scope string `json:"scope"`
	jwt.StandardClaims
}

const permissionKey = "permissionKey"

func (a *Auth) HoldPermissionsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
		tokenString := authHeaderParts[1]
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			cert, err := a.getPemCert(token)
			if err != nil {
				return nil, xerrors.Errorf("failed to getPemCert: %w", err)
			}
			result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return nil, xerrors.Errorf("failed to ParseRSAPublicKeyFromPEM: %w", err)
			}
			return result, nil
		})
		if err != nil {
			log.Printf("failed to ParseWithClaims: %w", err)
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if ok && token.Valid {
			permissions := strings.Split(claims.Scope, " ")
			for _, permission := range permissions {
				fmt.Println(permission)
			}
			h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), permissionKey, permissions)))
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
