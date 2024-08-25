package middlewares

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CabinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
		err = e.LoadPolicy()

		if err != nil {
			panic(err)
		}

		user := c.MustGet("role").(string) // Get user from context
		obj := c.Request.URL.Path          // Resource being accessed
		act := c.Request.Method            // HTTP method (GET, POST, etc.

		allowed, err := e.Enforce(user, obj, act)
		if err != nil || !allowed {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
