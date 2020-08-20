package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/alexsasharegan/dotenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/pr0f3ss0r-b/todolist-go-vue/api"
	"github.com/pr0f3ss0r-b/todolist-go-vue/database"
)

func main() {
	// Load .env variables
	if err := dotenv.Load(); err != nil {
		log.Println("Failed to load .env file")
	}

	// Database
	var err error
	database.DBCon, err = sql.Open("mysql", os.Getenv("DB_SOURCE"))
	if err != nil {
		panic(err.Error())
	}
	defer database.DBCon.Close()

	// API
	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.Static("../client/dist"))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "../client/dist",
		HTML5: true,
	}))

	apiGroup := e.Group("/api")
	api.ApiHandler(apiGroup)

	e.Start(":" + os.Getenv("PORT"))
}
