package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/Wojbeg/go-gorilla-crm/config"
	"github.com/Wojbeg/go-gorilla-crm/database"
	"github.com/Wojbeg/go-gorilla-crm/migration"
	"github.com/Wojbeg/go-gorilla-crm/routes"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rep := database.CreateInfosRepository(&config)

	migration.MigrateDatabase(rep)

	r := mux.NewRouter()

	router := routes.CreateNewInfoRouter(r, rep)
	router.Init()

	staticRouter := routes.CreateNewStaticRouter(r, rep)
	staticRouter.Init()

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0%s", config.Port),
		WriteTimeout: 90 * time.Second,
		ReadTimeout:  90 * time.Second,
	}

	log.Printf("Server started on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

