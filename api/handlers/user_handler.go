package api

import (
	"context"
	"github.com/Meenachinmay/hotel-reservation-golang/types"

	"github.com/Meenachinmay/hotel-reservation-golang/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := context.Background()
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams

	// get the data from the request
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	// validate the data got from the request
	if err := params.Validate(); err != nil {
		return err
	}

	// hash the password
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	// finally create the user and return it
	insertedUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}
