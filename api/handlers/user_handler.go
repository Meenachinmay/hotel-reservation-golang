package api

import (
	"github.com/Meenachinmay/hotel-reservation-golang/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Chinmay",
		LastName: "anand",
	}

	return c.JSON(u)

}