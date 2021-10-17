package repo

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

/*
	Hàm handle các routes
*/
func setUpRoutes(app *fiber.App) {
	app.Get("/users", listUsers)
	app.Post("/users", createUsers)
	app.Get("/users/id", getUser)
	app.Put("/users/id", updateUser)
	app.Delete("/users/id", deleteUser)
}

/*
	Hàm quản lý user repository với middleware ghi log request
*/
func RepoUsersWithMiddleware(port string) {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format: "[${method}] - ${path}\n",
		//Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":" + port))
}
