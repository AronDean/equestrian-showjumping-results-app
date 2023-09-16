package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"backend/utils"
)

type CtxKey string

func GetUserID(r *http.Request) int {
	return r.Context().Value(CtxKey("user_id")).(int)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.AuthorizationError(w, errors.New("authorization header is required"))
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		userID, err := utils.JwtDecode(os.Getenv("JWT_SECRET"), tokenString)
		if err != nil {
			utils.AuthorizationError(w, err)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), CtxKey("user_id"), userID))
		next.ServeHTTP(w, r)
	})
}
