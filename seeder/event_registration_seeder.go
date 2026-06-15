package seeder

import (
	"encoding/json"
	"log"
	"os"

	"gorm.io/gorm"

	"event-backend/entities"
)

type EventRegistrationSeeder struct{}

func NewEventRegistrationSeeder() *EventRegistrationSeeder {
	return &EventRegistrationSeeder{}
}

func (s *EventRegistrationSeeder) Handle(db *gorm.DB) error {
	data, err := os.ReadFile("seeder/files/event_registrations.json")
	if err != nil {
		return err
	}

	var rows []struct {
		UserEmail  string `json:"user_email"`
		EventSlug  string `json:"event_slug"`
		TicketType int    `json:"ticket_type"`
		Status     int    `json:"status"`
	}
	if err := json.Unmarshal(data, &rows); err != nil {
		return err
	}

	if err := db.Where("1 = 1").Delete(&entities.EventRegistrationEntity{}).Error; err != nil {
		return err
	}

	inserted := 0
	for _, row := range rows {
		var user entities.UserEntity
		if err := db.Where("email = ?", row.UserEmail).First(&user).Error; err != nil {
			log.Printf("EventRegistrationSeeder: user %q not found, skipping registration", row.UserEmail)
			continue
		}

		var event entities.EventEntity
		if err := db.Where("slug = ?", row.EventSlug).First(&event).Error; err != nil {
			log.Printf("EventRegistrationSeeder: event %q not found, skipping registration for %q", row.EventSlug, row.UserEmail)
			continue
		}

		var ticket entities.EventTicketEntity
		if err := db.Where("event_id = ? AND type = ?", event.Id, row.TicketType).First(&ticket).Error; err != nil {
			log.Printf("EventRegistrationSeeder: ticket type %d for event %q not found, skipping registration for %q", row.TicketType, row.EventSlug, row.UserEmail)
			continue
		}

		registration := entities.EventRegistrationEntity{
			UserId:        user.Id,
			EventId:       event.Id,
			EventTicketId: ticket.Id,
			Status:        row.Status,
			CreatedBy:     &user.Id,
		}
		if err := db.Create(&registration).Error; err != nil {
			return err
		}
		inserted++
	}

	if err := syncTicketRegisteredCounts(db); err != nil {
		return err
	}

	log.Printf("EventRegistrationSeeder: inserted %d event registrations", inserted)
	return nil
}

func syncTicketRegisteredCounts(db *gorm.DB) error {
	if err := db.Model(&entities.EventTicketEntity{}).Where("1 = 1").Update("registered_count", 0).Error; err != nil {
		return err
	}

	return db.Exec(`
		UPDATE event_tickets
		SET registered_count = registration_counts.total
		FROM (
			SELECT event_ticket_id, COUNT(*) AS total
			FROM event_registrations
			WHERE deleted_at IS NULL AND status IN (1, 2)
			GROUP BY event_ticket_id
		) AS registration_counts
		WHERE event_tickets.id = registration_counts.event_ticket_id
	`).Error
}
