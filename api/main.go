package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/thecaptainprice/dictionary-app/backend/config"
	"github.com/thecaptainprice/dictionary-app/backend/handlers"
	"github.com/thecaptainprice/dictionary-app/backend/repositories"
	"github.com/thecaptainprice/dictionary-app/backend/routers"
	"github.com/thecaptainprice/dictionary-app/backend/services"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := config.ConnectToDB(cfg)

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}
	defer db.Close()

	// Create repositories
	wordRepo := repositories.NewWordRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Create services
	wordService := services.NewWordService(wordRepo)
	userService := services.NewUserService(userRepo)

	// Create handlers
	wordHandler := handlers.NewWordHandler(wordService)
	userHandler := handlers.NewUserHandler(userService)

	// Create router and server
	router := routers.NewRouter(wordHandler, userHandler)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	// Start server
	fmt.Printf("Listening on port %s...\n", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
