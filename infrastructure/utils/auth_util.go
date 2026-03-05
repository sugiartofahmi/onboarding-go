package utils

import (
	"event-backend/infrastructure/exceptions"

	infraConstants "event-backend/infrastructure/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAuthUser(c *gin.Context) JWTUser {
	value, exists := c.Get(infraConstants.AuthUserKey)
	if !exists {
		panic(*exceptions.UnauthenticatedException("unauthenticated"))
	}

	user, ok := value.(JWTUser)
	if !ok {
		panic(*exceptions.UnauthenticatedException("invalid auth user in context"))
	}

	return user
}

func GetAuthUserId(c *gin.Context) uuid.UUID {
	return GetAuthUser(c).Id
}

func GetAuthRoleId(c *gin.Context) uuid.UUID {
	return GetAuthUser(c).RoleId
}
