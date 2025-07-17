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
	}))

	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.Get("/hello", c.HelloWorld)
			v1.Get("/body", c.BodyParser)
			v1.Post("/hello", c.Hello)
			v1.Put("/search", c.Search)
			v1.Post("/valid", c.Valid)
			v1.Get("/fact/:number", c.Fact)
			v1.Post("/register", c.Register)

			dog := v1.Group("/dog")
			{
				dog.Get("", c.GetDogs)
				dog.Get("/filter", c.GetDog)
				dog.Get("/json", c.GetDogsJson)
				dog.Post("/", c.AddDog)
				dog.Put("/:id", c.UpdateDog)
				dog.Delete("/:id", c.RemoveDog)
				dog.Get("/docs/deleted", c.GetDeletedDocs)
				dog.Get("/filter50", c.GetDogsFilter50)
			}
			company := v1.Group("/company")
			{
				company.Get("", c.GetCompanies)
				company.Get("/:id", c.GetCompany)
				company.Post("/", c.AddCompany)
				company.Put("/:id", c.UpdateCompany)
				company.Delete("/:id", c.RemoveCompany)
			}
		}

		v3 := api.Group("/v3")
		{
			v3.Get("/tong", c.Ascii)
		}
	}
}
