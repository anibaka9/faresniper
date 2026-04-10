package main

import (
	"log"
	"net/http"

	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const Limit int64 = 10

type Server struct {
	db *databaseclient.Client
	// сюда же шаблоны если будешь парсить при старте
}

func main() {
	c, err := databaseclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	server := Server{
		db: &c,
	}

	r := chi.NewRouter()

	r.Use(middleware.Compress(5, "text/html", "text/css"))

	r.Use(middleware.Logger)

	r.Get("/", server.handleIndex)
	r.Get("/routes", server.handleRoutes)
	r.Get("/airlines", server.handleAirlines)
	r.Get("/airports", server.handleAirports)
	r.Get("/cities", server.handleCities)
	r.Get("/countries", server.handleCountries)
	r.Get("/countries/{country_id}", server.handleCountry)
	r.Get("/airlines/{airline_id}", server.handleAirline)
	r.Get("/airports/{airport_id}", server.handleAirport)
	r.Get("/cities/{city_id}", server.handleCity)
	r.Get("/routes/{route_id}", server.handleRoute)

	r.Handle("/tailwind/*", http.StripPrefix("/tailwind/", http.FileServer(http.Dir("./cmd/web/tailwind"))))

	http.ListenAndServe(":3000", r)
}
