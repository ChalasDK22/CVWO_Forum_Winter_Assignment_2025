package middleware

import (
	"errors"
	"net/http"
	"strings"

	"chalas.com/forum_project/API/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if len(header) == 0 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Token is empty"))
			return
		}

		authParts := strings.SplitN(header, " ", 2)
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Token is invalid"))
			return
		}

		tokenStr := authParts[1]
		claims, err := jwt.ValidateJWTToken(tokenStr, secretKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
