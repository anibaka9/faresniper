package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleCities(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/cities.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	countCities, err := s.db.Queries.CountCities(r.Context())
	if err != nil {
		serverError(w, "count cities", err)
		return
	}

	cities, err := s.db.Queries.GetCities(r.Context(), database.GetCitiesParams{
		Limit:  Limit,
		Offset: Limit * int64(page),
	})

	fmt.Println(cities[0])

	if err != nil {
		serverError(w, "get cities", err)
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
		serverError(w, "render template", err)
		return
	}
}
