package seeder

import (
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, seederCommands []string) error {
	if len(seederCommands) > 0 {
		listSeeders := map[string]Seeder{
			"RoleSeeder":              NewRoleSeeder(),
			"CategorySeeder":          NewCategorySeeder(),
			"UserSeeder":              NewUserSeeder(),
			"EventSeeder":             NewEventSeeder(),
			"EventTicketSeeder":       NewEventTicketSeeder(),
			"EventRegistrationSeeder": NewEventRegistrationSeeder(),
		}
		for _, name := range seederCommands {
			s, ok := listSeeders[name]
			if !ok {
				log.Printf("unknown seeder: %s", name)
				continue
			}
			if err := s.Handle(db); err != nil {
				return err
			}
		}
	} else {
		// truncate in reverse FK order before full re-seed
		db.Exec(`DELETE FROM event_registrations`)
		db.Exec(`DELETE FROM event_tickets`)
		db.Exec(`DELETE FROM events`)
		db.Exec(`DELETE FROM users`)
		db.Exec(`DELETE FROM categories`)
		db.Exec(`DELETE FROM roles`)

		for _, s := range []Seeder{
			NewRoleSeeder(),
			NewCategorySeeder(),
			NewUserSeeder(),
			NewEventSeeder(),
			NewEventTicketSeeder(),
			NewEventRegistrationSeeder(),
		} {
			if err := s.Handle(db); err != nil {
				return err
			}
		}
	}
	log.Println("Seeding completed!")
	return nil
}
