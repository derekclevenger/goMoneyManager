package app

import (
	"fmt"
	"context"
	"github.com/derekclevenger/accountapi/utils"
	"github.com/derekclevenger/accountapi/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func (next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuthRequired := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path

		for _, v := range noAuthRequired {
			if v == requestPath {
				next.ServeHTTP(w,r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = utils.Message(false, "Missing Auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		split := strings.Split(tokenHeader, " ")
		if len(split) != 2 {
			response = utils.Message(false, "Invalid/Malformed Auth Token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		tokenPart := split[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = utils.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}

		if !token.Valid {
			response = utils.Message(false, "Token is invalid")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-Type", "application/json")
			utils.Respond(w, response)
			return
		}
		userId := fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(r.Context(), "user", userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
