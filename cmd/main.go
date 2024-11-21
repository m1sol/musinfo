package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"musinfo/internal/db"
	"musinfo/internal/routes"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Env load error: %v\n", err)
	}

	e := echo.New()

	db.ConnectDatabase()

	routes.Song(e)

	e.Logger.Fatal(e.Start(os.Getenv("ECHO_HOST")))
}
