package auth

import (
	"fmt"
	"net/http"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	debug               bool
	credentialsOptional bool
}

func New(debug, credentialsOptional bool) *Auth {
	return &Auth{debug, credentialsOptional}
}

func (a *Auth) Handler(next http.Handler) http.Handler {
	middleware := jwtMiddleware.New(jwtMiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			fmt.Println(token)
			// FIXME:
			return nil, nil
		},
		SigningMethod:       jwt.SigningMethodRS256,
		Debug:               a.debug,
		CredentialsOptional: a.credentialsOptional,
	})
	return middleware.Handler(next)
}
