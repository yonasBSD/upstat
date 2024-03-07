package routes

import (
  "net/http"

  "github.com/gofiber/fiber/v2/middleware/filesystem"
  "github.com/gofiber/fiber/v2"
)

// @Group UI
func UIRoutes(app *fiber.App) {
  app.Use(filesystem.New(filesystem.Config{
    Root: http.Dir("./web/dist"),
  }))
}
