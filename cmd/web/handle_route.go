package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleRoute(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/route.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	routeId, err := strconv.Atoi(chi.URLParam(r, "route_id"))
	if err != nil {
		http.Error(w, "wrong route id", http.StatusBadRequest)
		return
	}

	route, err := s.db.Queries.GetRoute(r.Context(), int64(routeId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		Route       database.GetRouteRow
	}{
		SidebarData: GetSidebarData("/routes"),
		Route:       route,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
