package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ivanauliaa/go-appoinment/src/api/echo/server"
	"github.com/ivanauliaa/go-appoinment/src/database"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env file")
	}

	database.InitMigration()
}

func main() {
	e := server.CreateServer()

	if err := e.Start(os.Getenv("PORT")); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
