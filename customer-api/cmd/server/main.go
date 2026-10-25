package main

import (
	"context"
	"log"

	"company/customer-api/internal/audit"
	"company/customer-api/internal/auth"
	"company/customer-api/internal/contact"
	"company/customer-api/internal/customer"
	"company/customer-api/internal/routes"
	"company/customer-api/internal/user"

	"company/customer-api/pkg/config"
	"company/customer-api/pkg/database"
	"company/customer-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	// --------------------------------------------------
	// Logger
	// --------------------------------------------------

	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}

	logger.Log.Info(
		"application_starting",
	)

	// --------------------------------------------------
	// Configuration
	// --------------------------------------------------

	cfg := config.Load()

	// --------------------------------------------------
	// Database
	// --------------------------------------------------

	db, err := database.NewPostgres(cfg)

	if err != nil {

		logger.Log.Fatal(
			"database_connection_failed",
		)
	}

	// --------------------------------------------------
	// Audit Module
	// --------------------------------------------------
	auditRepo := audit.NewRepository(
		db,
	)

	auditService := audit.NewService(
		auditRepo,
	)

	auditHandler := audit.NewHandler(
		auditService,
	)

	// --------------------------------------------------
	// User Module
	// --------------------------------------------------

	userRepo := user.NewRepository(
		db,
	)

	userService := user.NewService(
		userRepo, auditService,
	)

	userHandler := user.NewHandler(
		userService,
	)

	// --------------------------------------------------
	// Auth Module
	// --------------------------------------------------

	authService := auth.NewService(
		userRepo,
		cfg.JWTSecret,
		auditService,
	)

	authHandler := auth.NewHandler(
		authService,
	)
	// --------------------------------------------------
	// Seed Admin
	// --------------------------------------------------

	if err := auth.SeedAdmin(
		context.Background(),
		userService,
	); err != nil {

		logger.Log.Fatal(
			"admin_seed_failed",
		)
	}
	// --------------------------------------------------
	// Customer Module
	// --------------------------------------------------

	customerRepo := customer.NewPostgresRepository(
		db,
	)

	customerService := customer.NewService(
		customerRepo, auditService,
	)

	customerHandler := customer.NewHandler(
		customerService,
	)

	// --------------------------------------------------
	// Contact Module
	// --------------------------------------------------

	contactRepo := contact.NewRepository(
		db,
	)

	contactService := contact.NewService(
		contactRepo, auditService,
	)

	contactHandler := contact.NewHandler(
		contactService,
	)

	// --------------------------------------------------
	// Gin
	// --------------------------------------------------
	gin.SetMode(
		cfg.GIN_MODE,
	)
	r := gin.New()

	r.Use(
		logger.GinMiddleware(),
	)

	r.Use(
		gin.Recovery(),
	)

	// --------------------------------------------------
	// Routes
	// --------------------------------------------------

	routes.Register(
		r,
		cfg.JWTSecret,
		routes.Handlers{
			Auth:     authHandler,
			User:     userHandler,
			Customer: customerHandler,
			Contact:  contactHandler,
			Audit:    auditHandler,
		},
	)

	// --------------------------------------------------
	// Start Server
	// --------------------------------------------------

	logger.Log.Info(
		"server_started",
	)

	log.Fatal(
		r.Run(
			":" + cfg.Port,
		),
	)
}
