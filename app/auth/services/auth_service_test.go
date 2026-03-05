package services

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"event-backend/app/auth/mocks"
	"event-backend/app/role/constants"
	roleMocks "event-backend/app/role/mocks"
	"event-backend/entities"
	authdtos "event-backend/presentation/http/auth/dtos"
)

func TestAuthService_Login_Success(t *testing.T) {
	mockAuthQuery := &mocks.AuthQueryRepositoryMock{}
	mockAuthStore := &mocks.AuthStoreRepositoryMock{}
	mockRoleQuery := &roleMocks.RoleQueryRepositoryMock{}

	service := NewAuthService(mockAuthQuery, mockAuthStore, mockRoleQuery)

	userID := uuid.New()
	roleID := uuid.New()
	request := &authdtos.AuthLoginRequestDto{
		Email:    "test@example.com",
		Password: "password123",
	}

	userEntity := &entities.UserEntity{
		Id:        userID,
		RoleId:    roleID,
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  "$2a$10$hashedpassword",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockAuthQuery.On("FindOneByEmailWithRole", context.Background(), request.Email).Return(userEntity)

	result := service.Login(context.Background(), request)

	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, "Bearer", result.TokenType)
	mockAuthQuery.AssertExpectations(t)
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	mockAuthQuery := &mocks.AuthQueryRepositoryMock{}
	mockAuthStore := &mocks.AuthStoreRepositoryMock{}
	mockRoleQuery := &roleMocks.RoleQueryRepositoryMock{}

	service := NewAuthService(mockAuthQuery, mockAuthStore, mockRoleQuery)

	request := &authdtos.AuthLoginRequestDto{
		Email:    "notfound@example.com",
		Password: "password123",
	}

	mockAuthQuery.On("FindOneByEmailWithRole", context.Background(), request.Email).Return(nil)

	assert.Panics(t, func() {
		service.Login(context.Background(), request)
	})

	mockAuthQuery.AssertExpectations(t)
}

func TestAuthService_Register_Success(t *testing.T) {
	mockAuthQuery := &mocks.AuthQueryRepositoryMock{}
	mockAuthStore := &mocks.AuthStoreRepositoryMock{}
	mockRoleQuery := &roleMocks.RoleQueryRepositoryMock{}

	service := NewAuthService(mockAuthQuery, mockAuthStore, mockRoleQuery)

	roleID := uuid.New()
	request := &authdtos.AuthRegisterRequestDto{
		Name:     "New User",
		Email:    "newuser@example.com",
		Password: "password123",
	}

	roleEntity := &entities.RoleEntity{
		Id:        roleID,
		Name:      constants.ATTENDEE,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userEntity := &entities.UserEntity{
		Id:        uuid.New(),
		RoleId:    roleID,
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockAuthQuery.On("IsExistsByEmail", context.Background(), request.Email).Return(false)
	mockRoleQuery.On("FindOneByName", context.Background(), constants.ATTENDEE).Return(roleEntity)
	mockAuthStore.On("Create", context.Background(), mock.AnythingOfType("*entities.UserEntity")).Return(userEntity)

	result := service.Register(context.Background(), request)

	assert.NotNil(t, result)
	assert.Equal(t, request.Name, result.Name)
	assert.Equal(t, request.Email, result.Email)
	mockAuthQuery.AssertExpectations(t)
	mockRoleQuery.AssertExpectations(t)
	mockAuthStore.AssertExpectations(t)
}

func TestAuthService_Register_EmailAlreadyExists(t *testing.T) {
	mockAuthQuery := &mocks.AuthQueryRepositoryMock{}
	mockAuthStore := &mocks.AuthStoreRepositoryMock{}
	mockRoleQuery := &roleMocks.RoleQueryRepositoryMock{}

	service := NewAuthService(mockAuthQuery, mockAuthStore, mockRoleQuery)

	request := &authdtos.AuthRegisterRequestDto{
		Name:     "Existing User",
		Email:    "existing@example.com",
		Password: "password123",
	}

	mockAuthQuery.On("IsExistsByEmail", context.Background(), request.Email).Return(true)

	assert.Panics(t, func() {
		service.Register(context.Background(), request)
	})

	mockAuthQuery.AssertExpectations(t)
}

func TestAuthService_Register_RoleNotFound(t *testing.T) {
	mockAuthQuery := &mocks.AuthQueryRepositoryMock{}
	mockAuthStore := &mocks.AuthStoreRepositoryMock{}
	mockRoleQuery := &roleMocks.RoleQueryRepositoryMock{}

	service := NewAuthService(mockAuthQuery, mockAuthStore, mockRoleQuery)

	request := &authdtos.AuthRegisterRequestDto{
		Name:     "New User",
		Email:    "newuser@example.com",
		Password: "password123",
	}

	mockAuthQuery.On("IsExistsByEmail", context.Background(), request.Email).Return(false)
	mockRoleQuery.On("FindOneByName", context.Background(), constants.ATTENDEE).Return(nil)

	assert.Panics(t, func() {
		service.Register(context.Background(), request)
	})

	mockAuthQuery.AssertExpectations(t)
	mockRoleQuery.AssertExpectations(t)
}
