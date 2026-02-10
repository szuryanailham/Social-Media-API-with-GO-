package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/szuryanailham/social/internal/db"
	"github.com/szuryanailham/social/internal/env"
	storepkg "github.com/szuryanailham/social/internal/env/store"
)

const version = "0.0.1"
func main() {
	// Load .env file (optional)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using OS environment variables")
	}

	cfg := Config{
		Addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:admin12345@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetDuration("DB_MAX_IDLE_TIME", 15*time.Minute),
		},
		env: env.GetString("ENV","development"),
	}

	// Initialize database
	database, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime.String(), 
	)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	log.Println("Database connected pool established..")

	// Initialize store
	storage := storepkg.NewStorage(database)

	app := &Application{
		config: cfg,
		store:  storage,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
