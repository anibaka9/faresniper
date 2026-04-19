package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirportUpdatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/airport_update.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	airportId, err := strconv.Atoi(chi.URLParam(r, "airport_id"))
	if err != nil {
		http.Error(w, "wrong airport id", http.StatusBadRequest)
		return
	}

	airport, err := s.db.Queries.GetAirport(r.Context(), int64(airportId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cities, err := s.db.Queries.GetAllCities(r.Context())

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		Airport     database.GetAirportRow
		Cities      []database.GetAllCitiesRow
	}{
		SidebarData: GetSidebarData("/airports"),
		Airport:     airport,
		Cities:      cities,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
