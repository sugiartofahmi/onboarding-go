package validators

import (
	"github.com/google/uuid"

	"event-backend/infrastructure/exceptions"
)

func ValidateUUID(id string) uuid.UUID {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		panic(*exceptions.BadRequestException("Invalid UUID format"))
	}
	return parsedUUID
}

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
