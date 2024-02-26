package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sahaduta/backend2024-test-laravel/database"
	"github.com/sahaduta/backend2024-test-laravel/server"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := database.NewConn()
	if err != nil {
		log.Fatal("fail to connect to database")
	}

	r := server.NewRouter(server.GetRouterOpts(db))

	srv := http.Server{
		Addr:    os.Getenv("APP_PORT"),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Print("listen: %s\n", err)
	}

}
