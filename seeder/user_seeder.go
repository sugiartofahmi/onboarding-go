package seeder

import (
	"encoding/json"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
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

		hashed := hashPassword(row.Password)

		user := entities.UserEntity{
			RoleId:   role.Id,
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

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
