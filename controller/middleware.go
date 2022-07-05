package controller

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/arshabbir/bankapp/config"
	"github.com/arshabbir/bankapp/domain"
	"github.com/dgrijalva/jwt-go"
)

func (c *controller) AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Authrization middleware Invoked.")
		authStr := r.Header["Authorization"]
		log.Println("Auth string ", authStr, len(authStr))
		if len(authStr) == 0 {
			// no auth header
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message" : "no token"}`))
			return
		}
		tok := strings.Fields(authStr[0])

		if err := validateToken(tok[1]); err != nil {
			log.Println("error validating the token ", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message" : "invalied token"}`))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func validateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&domain.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GlobalConf.JWTKEY), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*domain.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
