package main

import (
	"log"
	"net/http"

	_ "go-lessons/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := chi.NewRouter()

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/docs/swagger.json"), //The url pointing to API definition
	))

	log.Fatal(http.ListenAndServe(":1323", r))
}