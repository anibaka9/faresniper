package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleRouteUpdatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/route_update.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	routeId, err := strconv.Atoi(chi.URLParam(r, "route_id"))
	if err != nil {
		http.Error(w, "wrong county id", http.StatusBadRequest)
		return
	}

	route, err := s.db.Queries.GetRoute(r.Context(), int64(routeId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	airports, err := s.db.Queries.GetAllAirports(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	airlines, err := s.db.Queries.GetAllAirlines(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		Route       database.GetRouteRow
		Airports    []database.GetAllAirportsRow
		Airlines    []database.GetAllAirlinesRow
	}{
		SidebarData: GetSidebarData("/routes"),
		Route:       route,
		Airports:    airports,
		Airlines:    airlines,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
