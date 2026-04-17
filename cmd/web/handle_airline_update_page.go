package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleAirlineUpdatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/airline_update.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	airline_id, err := strconv.Atoi(chi.URLParam(r, "airline_id"))
	if err != nil {
		http.Error(w, "wrong county id", http.StatusBadRequest)
		return
	}

	airline, err := s.db.Queries.GetAirline(r.Context(), int64(airline_id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
