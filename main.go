package main

import (
  "log"
  "os"

  "github.com/chamanbravo/upstat/database"
  _ "github.com/chamanbravo/upstat/docs"
  "github.com/chamanbravo/upstat/utils"

  "github.com/chamanbravo/upstat/routes"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger"
  _ "github.com/joho/godotenv/autoload"

  "github.com/ansrivas/fiberprometheus/v2"
)

// @title Upstat API
// @version 1.0
// @description This is an auto-generated API Docs for Upstat API.
// @contact.email chamanpro9@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
  // Start a new GoFiber application
  app := fiber.New()
  app.Use(logger.New())

  // Connect to the database
  if err := database.DBConnect(); err != nil {
    log.Fatal("Could not connect to database", err)
  }

  utils.StartGoroutineSetup()

  // Setup routes
  routes.AuthRoutes(app)
  routes.SwaggerRoute(app)
  routes.MonitorRoutes(app)
  routes.UserRoutes(app)
  routes.NotificationRoutes(app)
  routes.StatusPagesRoutes(app)
  routes.UIRoutes(app)

  // Export prometheus metrics
  prometheus := fiberprometheus.New("upstat")
  prometheus.RegisterAt(app, "/metrics")
  app.Use(prometheus.Middleware)

  // Get address and port
  port, ok := os.LookupEnv("PORT")
  if !ok {
    port = ":8000"
  } else {
    port = ":" + port
  }

  // Launch the server
  log.Fatal(app.Listen(port))
}
