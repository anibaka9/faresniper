package main

import (
	"net/http"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleWatchedRouteCreate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/watched_route_create.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	airports, err := s.db.Queries.GetAllAirports(r.Context())
	if err != nil {
		serverError(w, "get all airports", err)
		return
	}

	err = t.Execute(w, struct {
		SidebarData    []SidebarItem
		PaginationData PaginationData
		Airports       []database.GetAllAirportsRow
	}{
		SidebarData: GetSidebarData("/watchedroutes"),
		Airports:    airports,
	})
}
