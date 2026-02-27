package seeder

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"

	"event-backend/entities"
)

type UserSeeder struct{}

func NewUserSeeder() *UserSeeder {
	return &UserSeeder{}
}

func (s *UserSeeder) Handle(db *gorm.DB) error {
	data, err := os.ReadFile("seeder/files/users.json")
	if err != nil {
		return err
	}

	var rows []struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleName string `json:"role_name"`
	}
	if err := json.Unmarshal(data, &rows); err != nil {
		return err
	}

	if err := db.Where("1 = 1").Delete(&entities.UserEntity{}).Error; err != nil {
		return err
	}

	for _, row := range rows {
		var role entities.RoleEntity
		if err := db.Where("name = ?", row.RoleName).First(&role).Error; err != nil {
			log.Printf("UserSeeder: role %q not found, skipping user %q", row.RoleName, row.Email)
			continue
		}

		hashed, err := hashPassword(row.Password)
		if err != nil {
			return err
		}

		user := entities.UserEntity{
			RoleID:   role.ID,
			Name:     row.Name,
			Email:    row.Email,
			Password: hashed,
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	log.Printf("UserSeeder: inserted %d users", len(rows))
	return nil
}

func hashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("$argon2id$v=19$m=65536,t=1,p=4$%s$%s", encodedSalt, encodedHash), nil
}
