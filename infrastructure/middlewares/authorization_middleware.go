package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	infraConstants "event-backend/infrastructure/constants"
	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			panic(*exceptions.UnauthenticatedException("missing or invalid authorization header"))
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims := utils.ValidateToken(token)
		c.Set(infraConstants.AuthUserKey, claims.User)
		c.Next()
	}
}
