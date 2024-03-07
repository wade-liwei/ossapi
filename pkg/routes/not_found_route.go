// ./pkg/routes/not_found_route.go

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wade-liwei/ossapi/app"
	//github.com/wade-liwei/ossapi
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// upload files
	route.Post("/upload/privatedownload", app.PrivateDownloadFile)
	//route.Post("/upload/publicdownload", app.)
}
