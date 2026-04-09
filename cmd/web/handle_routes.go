package main

import (
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleRoutes(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/routes.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	countRoutes, err := s.db.Queries.CountRoutes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	routes, err := s.db.Queries.GetRoutes(r.Context(), database.GetRoutesParams{
		Limit:  Limit,
		Offset: Limit * int64(page),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paginationData := GetPaginationData(PaginationParams{
		TotalPages:   int64(math.Ceil(float64(countRoutes) / float64(Limit))),
		TotalCount:   countRoutes,
		CurrentPage:  int64(page),
		CurrentLimit: Limit,
	})

	err = t.Execute(w, struct {
		SidebarData    []SidebarItem
		PaginationData PaginationData
		Routes         []database.GetRoutesRow
	}{
		SidebarData:    GetSidebarData("/routes"),
		PaginationData: paginationData,
		Routes:         routes,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
