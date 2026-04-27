package main

import (
	"net/http"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirline(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/airline.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	iata := chi.URLParam(r, "airline_iata")
	if iata == "" {
		http.Error(w, "wrong airline iata", http.StatusBadRequest)
		return
	}

	airline, err := s.db.Queries.GetAirline(r.Context(), iata)
	if err != nil {
		serverError(w, "get airline "+iata, err)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		Airline     database.Airline
	}{
		SidebarData: GetSidebarData("/airlines"),
		Airline:     airline,
	})
	if err != nil {
		serverError(w, "render template", err)
		return
	}
}
