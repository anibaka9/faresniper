package main

import (
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleCities(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/cities.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	countCities, err := s.db.Queries.CountCities(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cities, err := s.db.Queries.GetCities(r.Context(), database.GetCitiesParams{
		Limit:  Limit,
		Offset: Limit * int64(page),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paginationData := GetPaginationData(PaginationParams{
		TotalPages:   int64(math.Ceil(float64(countCities) / float64(Limit))),
		TotalCount:   countCities,
		CurrentPage:  int64(page),
		CurrentLimit: Limit,
	})

	err = t.Execute(w, struct {
		SidebarData    []SidebarItem
		PaginationData PaginationData
		Cities         []database.GetCitiesRow
	}{
		SidebarData:    GetSidebarData("/cities"),
		PaginationData: paginationData,
		Cities:         cities,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
