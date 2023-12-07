package jwt_service

import (
	"context"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-blog-api-golang/configs"
	"simple-blog-api-golang/internal/user"
	"strings"
	"time"
)

func GenerateJWT(userModel *user.User) (string, error) {
	expiredAt := time.Now().Add(configs.LOGIN_EXPIRATION_DURATION)

	userClaims := user.NewUserClaims(jwt.StandardClaims{
		ExpiresAt: expiredAt.Unix(),
	}, userModel)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(configs.JWT_SIGNATURE_KEY)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Split the Authorization header into two parts: "Bearer" and the token
		tokenParts := strings.Fields(authorizationHeader)

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return configs.JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			fmt.Println("Error here")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next(w, r.WithContext(ctx), params)
	}
}
