package guards

import (
	"slices"

	"github.com/gin-gonic/gin"

	"event-backend/infrastructure/exceptions"
	"event-backend/infrastructure/utils"
)

func RoleGuard(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := utils.GetAuthUser(c)
		isRoleExists := slices.Contains(roles, user.RoleName)
		if !isRoleExists {
			panic(*exceptions.ForbiddenException("access denied"))
		}
		c.Next()
	}
}
