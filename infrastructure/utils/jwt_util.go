package utils

import (
	"event-backend/entities"
	"event-backend/infrastructure/config"
	infraConstants "event-backend/infrastructure/constants"
	"event-backend/infrastructure/exceptions"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTUser struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	RoleId   uuid.UUID `json:"role_id"`
	RoleName string    `json:"role_name"`
}

type JWTClaims struct {
	jwt.RegisteredClaims
	User JWTUser `json:"user"`
}

func GenerateToken(user *entities.UserEntity) (string, time.Time) {
	expiresIn, err := time.ParseDuration(config.JWTExpiredIn)
	if err != nil {
		panic(*exceptions.ServerErrorException(err))
	}
	expiresAt := time.Now().Add(expiresIn)

	claims := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		User: JWTUser{
			Id:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			RoleId:   user.RoleId,
			RoleName: user.Role.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		panic(*exceptions.ServerErrorException(err))
	}

	return tokenString, expiresAt
}

func ValidateToken(tokenString string) *JWTClaims {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		panic(*exceptions.UnauthenticatedException("invalid token"))
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		panic(*exceptions.UnauthenticatedException("invalid token claims"))
	}

	return claims
}

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
