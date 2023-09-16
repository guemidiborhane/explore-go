package auth

import "github.com/gofiber/fiber/v2"

var router fiber.Router

func setupRoutes() {
	group := router.Group("/auth")

	group.Get("/", CheckAuthenticated, Show)
	group.Post("/", validateUser, Register)
	group.Post("/session", validateLogin, Login)
	group.Delete("/session", Logout)
}
