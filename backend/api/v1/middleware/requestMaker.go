package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/casbin/casbin"
	jwt "github.com/dgrijalva/jwt-go"
	v1 "github.com/infinity-framework/backend/api/v1"
	"github.com/infinity-framework/backend/configs"
)

//RequestMiddleware Struct
type RequestMiddleware struct {
	Enforcer *casbin.Enforcer
}

//EnforcerFiles Func
func (rm RequestMiddleware) EnforcerFiles() *casbin.Enforcer {

	//Setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcerSafe("./configs/authPolicy/auth_model.conf", "./configs/authPolicy/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	return authEnforcer
}

//RequestIDGenerator Middleware
func (rm RequestMiddleware) RequestIDGenerator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(configs.WithRequestID(r.Context()))
		next.ServeHTTP(w, r)
	})
}

//ValidateMiddleware to authenticate jwt
func (rm RequestMiddleware) ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(configs.Config.JWTSecretKey), nil
				})
				if error != nil {
					v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid token")
					return
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					w.Header().Set("Access-Control-Allow-Credentials", "true")
					w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
					w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,x-requested-with, XMLHttpRequest, Access-Control-Allow-Methods")
					userType := claims["userType"].(string)
					enforce := rm.EnforcerFiles()
					// casbin rule enforcing
					res, err := enforce.EnforceSafe(userType, r.URL.Path, r.Method)
					if err != nil {
						v1.WriteErrorResponse(w, http.StatusInternalServerError, "Oops! Something went")
						return
					}
					if res {
						next.ServeHTTP(w, r)
					} else {
						v1.WriteErrorResponse(w, http.StatusForbidden, "Oops! You Can't Go There!")
						return
					}
				} else {
					v1.WriteErrorResponse(w, http.StatusBadRequest, "Invalid token")
				}
			}
		} else {
			v1.WriteErrorResponse(w, http.StatusBadRequest, "An Authorization Header is Required!")
		}
	})
}
