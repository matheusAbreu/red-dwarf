package main

import (
	"net/http"
	"github.com/matheusAbreu/red-dwarf/config"
	"github.com/matheusAbreu/red-dwarf/api/healthCheck"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/adaptor/v2"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Red Dwarf
// @version 1.0
// @description A Software for study goLang
// @contact.name Matheus Abreu
func main() {
    app := fiber.New(fiber.Config{
			// DisableStartupMessage: true,
		})

		if config.GetEnvBool("GO_PROFILING"){
			app.Use(pprof.New())
		}

		app.Use(compress.New())

		app.Use(cors.New())

		app.Get(healthCheck.PathRouter, healthCheck.Check)

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
		
		app.Get("/docs/", 
		adaptor.HTTPHandlerFunc(func(response http.ResponseWriter, request *http.Request){
			http.Redirect(response, request, "/docs/swagger/index.html", http.StatusMovedPermanently)
		}))

		app.Get("/docs/swagger/swagger.json", 
			adaptor.HTTPHandlerFunc(func(response http.ResponseWriter, request *http.Request){
				http.ServeFile(response, request, "./docs/swagger/swagger.json")
		}))

		app.Get("/docs/swagger/*", adaptor.HTTPHandlerFunc(httpSwagger.Handler(
			httpSwagger.URL("/docs/swagger/swagger.json"),
		)))

    app.Listen(":3000")

}