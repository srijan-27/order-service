package main

import (
	"gofr.dev/pkg/gofr"

	"github.com/srijan-27/order-service/handler"
	"github.com/srijan-27/order-service/migration"
	"github.com/srijan-27/order-service/service"
	"github.com/srijan-27/order-service/store"
)

func main() {
	// Create a new application
	app := gofr.New()

	// Run migrations
	app.Migrate(migration.All())

	s := store.New()
	svc := service.New(s)
	h := handler.New(svc)

	// Add required routes
	app.POST("/orders", h.Create)
	app.GET("/orders", h.GetAll)
	app.GET("/orders/{id}", h.GetByID)
	app.PUT("/orders/{id}", h.Update)
	app.DELETE("/orders/{id}", h.Delete)

	app.Run()
}
