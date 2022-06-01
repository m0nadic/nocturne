package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
)

func PrivateAuthMiddleware(signingKey string) gin.HandlerFunc {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "snippets",
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		Key:           []byte(signingKey),
	})

	if err != nil {
		log.Println("JWT Error")
	}

	return authMiddleware.MiddlewareFunc()
}
