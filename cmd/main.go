package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"pwd-monitor/internal/app/handlers"
	"pwd-monitor/internal/app/middleware"
	"pwd-monitor/internal/app/models"
	"pwd-monitor/pkg/db"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Connect to the database
	db, err := db.Connect()
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// Apply middleware to log requests and responses
	e.Use(middleware.Logger())

	// Apply authentication middleware
	e.Use(pwdmonitormiddleware.Authenticate())

	// Define routes for the home page, projects, site inspections, reports, and financial monitoring
	e.GET("/", handlers.Home)
	e.GET("/projects", handlers.Projects)
	e.GET("/site-inspections", handlers.SearchSiteInspections)
	e.GET("/reports", handlers.Reports)
	e.GET("/financial-monitoring", handlers.FinancialMonitoring)

	// Define routes for user registration and login
	e.GET("/register", handlers.Register)
	e.POST("/register", handlers.HandleRegistration)
	e.GET("/login", handlers.Login)
	e.POST("/login", handlers.HandleLogin)

	// Start the web server
	e.Logger.Fatal(e.Start(":8080"))
}
