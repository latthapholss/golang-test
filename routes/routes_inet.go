package routes

import (
	c "golang-training/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
		app.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				"gofiber": "21022566",
			},
			Unauthorized: func(c *fiber.Ctx) error {
				c.Set(fiber.HeaderWWWAuthenticate, `Basic realm="Restricted"`)
				return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
			},
		}))
		api := app.Group("/api") // /api
		{
			v1 := api.Group("/v1") // /api/v1
			{
				v1.Get("/hello", c.HelloWorld)
				v1.Get("/body", c.BodyParser)
				v1.Post("/hello", c.Hello)
				v1.Put("/search", c.Search)
				v1.Post("/valid", c.Valid)
				v1.Get("/fact/:number", c.Fact)
				v1.Post("register", c.Register)
			}
			v3 := api.Group("/v3") 

			{
				v3.Get("/tong", c.Ascii)
			}
		}

	}
