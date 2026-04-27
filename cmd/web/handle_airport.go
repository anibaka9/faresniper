package main

import (
	"net/http"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirport(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/airport.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	iata := chi.URLParam(r, "airport_id")
	if iata == "" {
		http.Error(w, "wrong airport id", http.StatusBadRequest)
		return
	}

	airport, err := s.db.Queries.GetAirport(r.Context(), iata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		Airport     database.GetAirportRow
	}{
		SidebarData: GetSidebarData("/airports"),
		Airport:     airport,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
