package seeder

import (
	"encoding/json"
	"log"
	"os"

	"gorm.io/gorm"

	"event-backend/entities"
)

type CategorySeeder struct{}

func NewCategorySeeder() *CategorySeeder {
	return &CategorySeeder{}
}

func (s *CategorySeeder) Handle(db *gorm.DB) error {
	data, err := os.ReadFile("seeder/files/categories.json")
	if err != nil {
		return err
	}

	var rows []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	if err := json.Unmarshal(data, &rows); err != nil {
		return err
	}

	if err := db.Where("1 = 1").Delete(&entities.CategoryEntity{}).Error; err != nil {
		return err
	}

	for _, row := range rows {
		category := entities.CategoryEntity{Name: row.Name, Slug: row.Slug}
		if err := db.Create(&category).Error; err != nil {
			return err
		}
	}

	log.Printf("CategorySeeder: inserted %d categories", len(rows))
	return nil
}
