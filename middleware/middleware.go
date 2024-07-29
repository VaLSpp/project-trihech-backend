package middleware

import (
	"net/http"
	"project-trihech-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized header required"})
            c.Abort()
            return
        }


        parts := strings.SplitN(tokenString, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
            c.Abort()
            return
        }

        // Contoh validasi token (disesuaikan dengan logika autentikasi Anda)
        tokenString = parts[1]
        claims, err := utils.ValidateJWT(tokenString)

        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
            c.Abort()
            return
        }

        c.Set("UserID", claims.UserID)
        c.Next()
    }
}
