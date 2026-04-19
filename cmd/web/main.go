package main

import (
	"log"
	"net/http"

	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mattn/go-sqlite3"
)

const Limit int64 = 10

var sqliteErr sqlite3.Error

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
	r.Get("/countries/{country_id}/edit", server.handleCountryUpdatePage)
	r.Post("/countries/{country_id}", server.handleCountryUpdate)
	r.Get("/airlines/{airline_id}", server.handleAirline)
	r.Get("/airlines/{airline_id}/edit", server.handleAirlineUpdatePage)
	r.Post("/airlines/{airline_id}", server.handleAirlineUpdate)
	r.Get("/airports/{airport_id}", server.handleAirport)
	r.Get("/airports/{airport_id}/edit", server.handleAirportUpdatePage)
	r.Post("/airports/{airport_id}", server.handleAirportUpdate)
	r.Get("/cities/{city_id}", server.handleCity)
	r.Get("/cities/{city_id}/edit", server.handleCityUpdatePage)
	r.Post("/cities/{city_id}", server.handleCityUpdate)
	r.Get("/routes/{route_id}", server.handleRoute)
	r.Get("/routes/{route_id}/edit", server.handleRouteUpdatePage)
	r.Post("/routes/{route_id}", server.handleRouteUpdate)

	r.Handle("/tailwind/*", http.StripPrefix("/tailwind/", http.FileServer(http.Dir("./cmd/web/tailwind"))))

	http.ListenAndServe(":3000", r)
}
