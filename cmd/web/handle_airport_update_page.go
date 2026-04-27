package main

import (
	"net/http"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirportUpdatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/airport_update.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	iata := chi.URLParam(r, "airport_iata")
	if iata == "" {
		http.Error(w, "wrong airport id", http.StatusBadRequest)
		return
	}

	airport, err := s.db.Queries.GetAirport(r.Context(), iata)
	if err != nil {
		serverError(w, "get airport "+iata, err)
		return
	}

	cities, err := s.db.Queries.GetAllCities(r.Context())
	if err != nil {
		serverError(w, "get cities", err)
		return
	}

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
		serverError(w, "render template", err)
		return
	}
}
