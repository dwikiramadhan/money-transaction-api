package main

import (
	"os"

	"github.com/dwikiramadhan/money-transaction-api/app/dal"
	"github.com/dwikiramadhan/money-transaction-api/app/dao"
	"github.com/dwikiramadhan/money-transaction-api/pkg/configs"
	"github.com/dwikiramadhan/money-transaction-api/pkg/middleware"
	"github.com/dwikiramadhan/money-transaction-api/pkg/routes"
	"github.com/dwikiramadhan/money-transaction-api/pkg/utils"
	"github.com/dwikiramadhan/money-transaction-api/platform/database"

	"github.com/gofiber/fiber/v2"

	_ "github.com/dwikiramadhan/money-transaction-api/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title Test 2 REST API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	database.OpenDBConnection()
	dao.DB = database.DB.Db

	dal.SetDefault(database.DB.Db)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
