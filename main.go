package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"       // DB user
	password = "Fikhri12"       // DB password
	dbname   = "pelabuhansigma" // DB name
)

var db *sql.DB

type Shipment struct {
	ID            int       `json:"id"`
	ShipID        int       `json:"ship_id"`
	CargoID       int       `json:"cargo_id"`
	TerminalID    int       `json:"terminal_id"`
	ArrivalTime   time.Time `json:"arrival_time"`
	DepartureTime time.Time `json:"departure_time"`
	Status        string    `json:"status"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	port := 3000

	// API routes
	app.Get("/shipments/:id", getShipmentByID)
	app.Post("/shipments", createShipment)

	log.Println("Fiber API running on ", port)
	log.Fatal(app.Listen(":" + fmt.Sprint(port)))
}

func getShipmentByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var shipment Shipment

	query := `
		SELECT id, ship_id, cargo_id, terminal_id, arrival_time, departure_time, status
		FROM shipments WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&shipment.ID, &shipment.ShipID, &shipment.CargoID,
		&shipment.TerminalID, &shipment.ArrivalTime, &shipment.DepartureTime, &shipment.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Shipment not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database query error",
		})
	}

	return c.JSON(shipment)
}

func createShipment(c *fiber.Ctx) error {
	shipment := new(Shipment)
	if err := c.BodyParser(shipment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	query := `
		INSERT INTO shipments (ship_id, cargo_id, terminal_id, arrival_time, departure_time, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err := db.QueryRow(query, shipment.ShipID, shipment.CargoID, shipment.TerminalID,
		shipment.ArrivalTime, shipment.DepartureTime, shipment.Status).Scan(&shipment.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create shipment",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(shipment)
}
