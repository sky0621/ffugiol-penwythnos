package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/sky0621/fs-mng-backend/src/util"

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

// context格納時のキー
const authenticatedUserKey = "authenticatedUserKey"

// 認証チェック済みのユーザー情報を保持
type AuthenticatedUser struct {
	ID            string
	EMail         string
	PermissionSet *util.StringSet
	// FIXME: 所属する organization の情報を別途保持
}

func (u *AuthenticatedUser) HasPermission(funcName string, crud operation, t target) bool {
	if u == nil || u.PermissionSet == nil {
		return false
	}
	return u.PermissionSet.Contains(fmt.Sprintf("%s:%v:%v", funcName, crud, t))
}

func (u *AuthenticatedUser) HasNoTargetPermission(funcName string, crud operation) bool {
	if u == nil || u.PermissionSet == nil {
		return false
	}
	return u.PermissionSet.Contains(fmt.Sprintf("%s:%v", funcName, crud))
}

func (u *AuthenticatedUser) HasReadAllPermission(funcName string) bool {
	return u.HasPermission(funcName, READ, ALL)
}

func (u *AuthenticatedUser) HasReadMinePermission(funcName string) bool {
	return u.HasPermission(funcName, READ, MINE)
}

func (u *AuthenticatedUser) HasCreatePermission(funcName string) bool {
	return u.HasNoTargetPermission(funcName, CREATE)
}

func (a *Auth) HoldPermissionsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
		tokenString := authHeaderParts[1]

		parser := func(token *jwt.Token) (interface{}, error) {
			cert, err := a.getPemCert(token)
			if err != nil {
				return nil, xerrors.Errorf("failed to getPemCert: %w", err)
			}
			result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return nil, xerrors.Errorf("failed to ParseRSAPublicKeyFromPEM: %w", err)
			}
			return result, nil
		}

		token, err := jwt.Parse(tokenString, parser)
		if err != nil {
			log.Printf("failed to ParseWithClaims: %v", err)
			next.ServeHTTP(w, r)
			return
		}
		if !token.Valid {
			log.Print("invalid token")
			next.ServeHTTP(w, r)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Print("not implemented MapClaims")
			next.ServeHTTP(w, r)
			return
		}

		authenticatedUser := &AuthenticatedUser{
			PermissionSet: util.NewBlankStringSet(),
		}
		for k, v := range claims {
			switch k {
			case "sub":
				authenticatedUser.ID = v.(string)
			case a.audience + "/email":
				authenticatedUser.EMail = v.(string)
			case "permissions":
				if permissionArray, ok := v.([]interface{}); ok {
					for _, permission := range permissionArray {
						authenticatedUser.PermissionSet.Add(permission.(string))
					}
					break
				}
			}
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), authenticatedUserKey, authenticatedUser)))
	})
}

func GetAuthenticatedUser(ctx context.Context) *AuthenticatedUser {
	u, ok := ctx.Value(authenticatedUserKey).(*AuthenticatedUser)
	if !ok {
		return nil
	}
	return u
}

type operation string
type target string

const (
	READ   operation = "read"
	CREATE           = "create"
	UPDATE           = "update"
	DELETE           = "delete"

	ALL  target = "all"
	MINE        = "mine"
)
