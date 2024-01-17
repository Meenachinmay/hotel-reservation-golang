package main

import (
	api "github.com/Meenachinmay/hotel-reservation-golang/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
	apiv1 := app.Group("/api/v1")

    apiv1.Get("/user", api.HandleGetUsers)

    app.Listen(":3000")
}
