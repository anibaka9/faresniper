package main

import (
	"fmt"
	"log"
	"net/http"

	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mattn/go-sqlite3"
)

const Limit int64 = 10

var sqliteErr sqlite3.Error

func serverError(w http.ResponseWriter, msg string, err error) {
	http.Error(w, fmt.Sprintf("%s: %v", msg, err), http.StatusInternalServerError)
}

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
	r.Get("/airlines", server.handleAirlines)
	r.Get("/airports", server.handleAirports)
	r.Get("/cities", server.handleCities)
	r.Get("/countries", server.handleCountries)
	r.Get("/countries/{country_code}", server.handleCountry)
	r.Get("/countries/{country_code}/edit", server.handleCountryUpdatePage)
	r.Post("/countries/{country_code}", server.handleCountryUpdate)
	r.Get("/airlines/{airline_iata}", server.handleAirline)
	r.Get("/airlines/{airline_iata}/edit", server.handleAirlineUpdatePage)
	r.Post("/airlines/{airline_iata}", server.handleAirlineUpdate)
	r.Get("/airports/{airport_iata}", server.handleAirport)
	r.Get("/airports/{airport_iata}/edit", server.handleAirportUpdatePage)
	r.Post("/airports/{airport_iata}", server.handleAirportUpdate)
	r.Get("/cities/{city_iata}", server.handleCity)
	r.Get("/cities/{city_iata}/edit", server.handleCityUpdatePage)
	r.Post("/cities/{city_iata}", server.handleCityUpdate)
	r.Get("/watchedroutes", server.handleWatchedRoutes)
	r.Get("/watchedroutes/new", server.handleWatchedRouteCreate)

	r.Handle("/tailwind/*", http.StripPrefix("/tailwind/", http.FileServer(http.Dir("./cmd/web/tailwind"))))

	http.ListenAndServe(":3000", r)
}
