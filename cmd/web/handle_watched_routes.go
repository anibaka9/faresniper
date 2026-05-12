package main

import (
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleWatchedRoutes(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/watched_routes.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	countWatchedRoutes, err := s.db.Queries.CountWatchedRoutes(r.Context())
	if err != nil {
		serverError(w, "cant get watched routes count", err)
		return
	}
	watchedRoutes, err := s.db.Queries.GetWatchedRoutes(r.Context(), database.GetWatchedRoutesParams{
		Limit:  Limit,
		Offset: Limit * int64(page),
	})
	if err != nil {
		serverError(w, "cant get watched routes", err)
		return
	}
	paginationData := GetPaginationData(PaginationParams{
		TotalPages:   int64(math.Ceil(float64(countWatchedRoutes) / float64(Limit))),
		TotalCount:   countWatchedRoutes,
		CurrentPage:  int64(page),
		CurrentLimit: Limit,
	})

	err = t.Execute(w, struct {
		SidebarData    []SidebarItem
		PaginationData PaginationData
		WatchedRoutes  []database.WatchedRoute
	}{
		SidebarData:    GetSidebarData("/watchedroutes"),
		PaginationData: paginationData,
		WatchedRoutes:  watchedRoutes,
	})
}
