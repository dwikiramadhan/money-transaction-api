package routes

import (
	"github.com/dwikiramadhan/money-transaction-api/app/controllers"
	"github.com/dwikiramadhan/money-transaction-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/")

	route.Post("/topup", middleware.JWTProtected(), controllers.Topup)
	route.Post("/pay", middleware.JWTProtected(), controllers.Payment)
	route.Post("/transfer", middleware.JWTProtected(), controllers.Transfer)
	route.Get("/transactions", middleware.JWTProtected(), controllers.Login)
	route.Put("/profile", middleware.JWTProtected(), controllers.Login)
}
