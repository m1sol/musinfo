package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"musinfo/docs"
	"musinfo/internal/db"
	"musinfo/internal/routes"
	"net/url"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Env load error: %v\n", err)
	}

	hostURL := os.Getenv("SWAGGER_HOST")
	// Разбираем URL
	parsedURL, err := url.Parse(hostURL)

	docs.SwaggerInfo.Title = "API онлайн библиотеки песен v.1.01"
	docs.SwaggerInfo.Description = "API онлайн библиотеки песен следует использовать только в develop режиме."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = parsedURL.Host

	e := echo.New()

	db.ConnectDatabase()

	routes.Song(e)

	e.Logger.Fatal(e.Start(os.Getenv("ECHO_HOST")))
}
