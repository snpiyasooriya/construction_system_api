package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuthentication(secretKey string) gin.HandlerFunc {
	fmt.Println(secretKey)
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("jwt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("error getting authorization token %s", err)})
			c.Abort()
			return
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization is missing"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			panic(err)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Set custom claims in context
			c.Set("user_id", claims["sub"])
			c.Set("role", claims["role"])
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
