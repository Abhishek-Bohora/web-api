// handler_product.go
package main

import (
	"net/http"
	"time"

	"github.com/Abhishek-Bohora/web-api/internal/database"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (app *App) handlerCreateProduct(c echo.Context) error {
	// Parse the request data from the client
	name := c.FormValue("name")
	userIDStr := c.FormValue("user_id")

	// Parse the userID string to a valid UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: "invalid user ID format",
		})
	}

	// Call the DB function to insert the product
	product, err := app.DB.CreateProduct(c.Request().Context(), database.CreateProductParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		UserID:    userID,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "couldn't create product",
		})
	}

	return c.JSON(http.StatusCreated, product)
}
