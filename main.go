package main

import (
	"os"
	"test_golang/controllers"
	"test_golang/initializers"

	"github.com/gofiber/fiber/v2"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	// SETUP App
	app := fiber.New()

	// CONFIGURE APP
	app.Static("/", "./public")

	// ROUTES
    app.Get("/", controllers.PostsIndex)
	// Start app
    app.Listen(":"+os.Getenv("PORT"))
}