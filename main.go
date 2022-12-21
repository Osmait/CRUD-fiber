package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/osmait/crud-fiber/database"
	"github.com/osmait/crud-fiber/lead"
)

func setRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("connection opened to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setRoutes(app)
	app.Listen(3000)
	fmt.Println("Server on port 3000")
	defer database.DBconn.Close()
}
