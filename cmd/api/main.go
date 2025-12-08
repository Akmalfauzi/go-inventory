package main

import (
	"context"
	"go-inventory/internal/db"
	"go-inventory/internal/delivery/http"
	"go-inventory/internal/repository/postgres"
	"go-inventory/internal/usecase"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Load Env
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is missing")
	}

	// 2. Connect DB Pool
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// 3. Setup Dependencies (Wiring Clean Arch)

	// Init DB Layer
	sqlcQueries := db.New(pool)

	// Init Repository (Layer 2)
	productRepo := postgres.NewProductRepository(sqlcQueries)

	// Init UseCase (Layer 3) - Set Global Timeout 2 detik
	productUseCase := usecase.NewProductUseCase(productRepo, 2*time.Second)

	// 4. Setup Fiber
	app := fiber.New(fiber.Config{AppName: "Go Inventory Clean Arch"})
	app.Use(logger.New())
	app.Use(cors.New())

	// Init Delivery (Layer 4)
	http.NewProductHandler(app, productUseCase)

	// 5. Start
	log.Println("🚀 Server running on port 3000")
	log.Fatal(app.Listen(":3000"))
}
