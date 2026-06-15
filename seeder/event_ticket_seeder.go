package seeder

import (
	"encoding/json"
	"log"
	"os"

	"gorm.io/gorm"

	"event-backend/entities"
)

type EventTicketSeeder struct{}

func NewEventTicketSeeder() *EventTicketSeeder {
	return &EventTicketSeeder{}
}

func (s *EventTicketSeeder) Handle(db *gorm.DB) error {
	data, err := os.ReadFile("seeder/files/event_tickets.json")
	if err != nil {
		return err
	}

	var rows []struct {
		EventSlug string  `json:"event_slug"`
		Type      int     `json:"type"`
		Price     float64 `json:"price"`
		Quota     int     `json:"quota"`
	}
	if err := json.Unmarshal(data, &rows); err != nil {
		return err
	}

	if err := db.Where("1 = 1").Delete(&entities.EventTicketEntity{}).Error; err != nil {
		return err
	}

	inserted := 0
	for _, row := range rows {
		var event entities.EventEntity
		if err := db.Where("slug = ?", row.EventSlug).First(&event).Error; err != nil {
			log.Printf("EventTicketSeeder: event %q not found, skipping ticket type %d", row.EventSlug, row.Type)
			continue
		}

		ticket := entities.EventTicketEntity{
			EventId: event.Id,
			Type:    row.Type,
			Price:   row.Price,
			Quota:   row.Quota,
		}
		if err := db.Create(&ticket).Error; err != nil {
			return err
		}
		inserted++
	}

	log.Printf("EventTicketSeeder: inserted %d event tickets", inserted)
	return nil
}
