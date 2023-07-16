package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Abhishek-Bohora/web-api/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type App struct {
	DB *database.Queries
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	// CORS middleware
	e.Use(middleware.CORS()) 

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("port number is not set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("database url is not set")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	app := &App{
		DB: database.New(conn),
	}

	//users
	e.POST("/user/create", app.handlerCreateUser)

	//products
	e.POST("user/product/create", app.handlerCreateProduct)

	// Start the server
	e.Logger.Fatal(e.Start(":1323"))
	
}


func (app *App) handlerCreateUser(c echo.Context) error {
	name := c.FormValue("name")

	user, err := app.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "couldn't create user",
		})
	}

	fmt.Println(user)
	return c.JSON(http.StatusCreated, user)
}