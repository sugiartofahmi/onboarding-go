package seeder

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"gorm.io/gorm"

	"event-backend/entities"
	"event-backend/infrastructure/utils"
)

type EventSeeder struct{}

func NewEventSeeder() *EventSeeder {
	return &EventSeeder{}
}

func (s *EventSeeder) Handle(db *gorm.DB) error {
	data, err := os.ReadFile("seeder/files/events.json")
	if err != nil {
		return err
	}

	var rows []struct {
		Title          string `json:"title"`
		CategorySlug   string `json:"category_slug"`
		OrganizerEmail string `json:"organizer_email"`
		Description    string `json:"description"`
		Location       string `json:"location"`
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
		Status         int    `json:"status"`
	}
	if err := json.Unmarshal(data, &rows); err != nil {
		return err
	}

	if err := db.Where("1 = 1").Delete(&entities.EventEntity{}).Error; err != nil {
		return err
	}

	inserted := 0
	for _, row := range rows {
		var category entities.CategoryEntity
		if err := db.Where("slug = ?", row.CategorySlug).First(&category).Error; err != nil {
			log.Printf("EventSeeder: category %q not found, skipping event %q", row.CategorySlug, row.Title)
			continue
		}

		var organizer entities.UserEntity
		if err := db.Where("email = ?", row.OrganizerEmail).First(&organizer).Error; err != nil {
			log.Printf("EventSeeder: organizer %q not found, skipping event %q", row.OrganizerEmail, row.Title)
			continue
		}

		startDate, err := time.Parse(time.RFC3339, row.StartDate)
		if err != nil {
			return err
		}

		endDate, err := time.Parse(time.RFC3339, row.EndDate)
		if err != nil {
			return err
		}

		event := entities.EventEntity{
			OrganizerUserId: organizer.Id,
			CategoryId:      category.Id,
			Title:           row.Title,
			Slug:            utils.GenerateSlug(row.Title),
			Description:     &row.Description,
			Location:        &row.Location,
			StartDate:       startDate,
			EndDate:         endDate,
			Status:          row.Status,
			CreatedBy:       &organizer.Id,
		}
		if err := db.Create(&event).Error; err != nil {
			return err
		}
		inserted++
	}

	log.Printf("EventSeeder: inserted %d events", inserted)
	return nil
}
