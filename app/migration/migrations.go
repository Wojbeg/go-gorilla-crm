package migration

import (
	"log"

	"github.com/Wojbeg/go-gorilla-crm/database"
	"github.com/Wojbeg/go-gorilla-crm/models"
)

func MigrateDatabase(rep database.Repository) {
	err := rep.AutoMigrate(&models.PersonInfo{})

	if err != nil {
		log.Fatal("Could not migrate person info: ", err)
	}
}