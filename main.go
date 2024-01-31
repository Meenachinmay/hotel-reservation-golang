package main

import (
	"context"
	"log"

	api "github.com/Meenachinmay/hotel-reservation-golang/api/handlers"
	"github.com/Meenachinmay/hotel-reservation-golang/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbURI = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const collection_name = "users"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)

	app.Listen(":3000")
}
