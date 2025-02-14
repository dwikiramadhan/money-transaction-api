package routes

import (
	"github.com/dwikiramadhan/money-transaction-api/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/")

	// Routes for POST method:
	route.Post("/register", controllers.Register) // register a new user
	route.Post("/login", controllers.Login)       // auth, return Access & Refresh tokens
}
