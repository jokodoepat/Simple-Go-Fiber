package main

import (
	"go-fiber-api/config"
	"go-fiber-api/database"
	"go-fiber-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

// type Movie struct {
//     ID    int    `json:"id"`
//     Title string `json:"title"`
//     Genre string `json:"genre"`
// }

// var movies = []Movie{
//     {ID: 1, Title: "The Matrix", Genre: "Action"},
//     {ID: 2, Title: "Inception", Genre: "Sci-Fi"},
//     {ID: 3, Title: "Titanic", Genre: "Drama"},
// }

func main() {

	// Inisialisasi koneksi database
	database.ConnectDB()

	// Buat instance Fiber atau Inisiasi fiber
	app := fiber.New()

	// Load configuration
	config.LoadConfig()

	// Setup routes
	routes.SetupRoutes(app)

	// app.Get("/api/v1/hello", func(c *fiber.Ctx) error {
	//     return c.JSON(fiber.Map{
	//         "message": "Hello, World!",
	//     })
	// })

	// // Route untuk mendapatkan semua film
	// app.Get("/movies", func(c *fiber.Ctx) error {
	//     return c.JSON(movies)
	// })

	// // Route untuk mendapatkan film berdasarkan ID
	// app.Get("/movies/:id", func(c *fiber.Ctx) error {
	//     id := c.Params("id")
	//     for _, movie := range movies {
	//         if id == string(movie.ID) {
	//             return c.JSON(movie)
	//         }
	//     }
	//     return c.Status(404).SendString("Movie not found")
	// })

	// // Route untuk endpoint POST
	// app.Post("/api/v1/greet", func(c *fiber.Ctx) error {
	//     type Request struct {
	//         Name string `json:"name"`
	//     }

	//     var req Request
	//     if err := c.BodyParser(&req); err != nil {
	//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//             "error": "Invalid input",
	//         })
	//     }

	//     return c.JSON(fiber.Map{
	//         "message": "Hello, " + req.Name + "!",
	//     })
	// })

	// Menjalankan server di port 3000
	app.Listen(":" + config.AppConfig.Port)
	// app.Listen(":3000")
}
