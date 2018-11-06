package authentication

import (
	"app/models"
	"app/utils"
	"errors"

	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

type JwtToken struct {
	Token string `json:"token"`
}

type AuthException struct {
	Message string `json:"message"`
}

var signedKey = []byte("MyAppSecretKey")

func GenerateJwtToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	utils.HandleError(err)
	signedToken, err := SignedJwtToken(&user)
	if err != nil {
		json.NewEncoder(w).Encode(AuthException{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(JwtToken{Token: signedToken})
}

func SignedJwtToken(user *models.UserLogin) (string, error) {
	// Create a new token object, specifying signing method and the claims  to be contained
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
	})

	// Sign and get the complete encoded token as a string using the secret
	signedToken, err := token.SignedString(signedKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

/* Middleware handler to handle all requests for authentication */
func AuthenticationMiddleware(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

		authHeader := req.Header.Get("Authorization")

		if authHeader != "" {

			bearerToken := strings.Split(authHeader, " ")

			if len(bearerToken) == 2 {

				token, err := jwt.Parse(bearerToken[1], lookupValidatingKey)
				// catch token errors
				if err != nil {
					utils.HandleError(err)
					w.WriteHeader(401)
					return
				}
				if token.Valid {
					utils.LogInfo("TOKEN WAS VALID")
					//Add data to context
					ctx := context.WithValue(req.Context(), "decoded", token.Claims)
					next(w, req.WithContext(ctx), ps)
				} else {
					utils.HandleError(errors.New("Invalid authentication token"))
					w.WriteHeader(401)
				}
			}
		} else {
			utils.HandleError(errors.New("An Authorization header is required"))
			w.WriteHeader(401)
		}
	}
}

func lookupValidatingKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("There was an error")
	}
	return signedKey, nil
}
