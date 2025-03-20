package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/hamza-s47/crm/database"
	"github.com/hamza-s47/crm/lead"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect Database")
	}
	log.Println("Connection opened to Database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	log.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(9020)
	defer database.DBConn.Close()
}
