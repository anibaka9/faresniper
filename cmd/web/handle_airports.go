package main

import (
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleAirports(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/airports.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	countAirports, err := s.db.Queries.CountAirports(r.Context())
	if err != nil {
		serverError(w, "count airports", err)
		return
	}

	airports, err := s.db.Queries.GetAirports(r.Context(), database.GetAirportsParams{
		Limit:  Limit,
		Offset: Limit * int64(page),
	})
	if err != nil {
		serverError(w, "get airports", err)
		return
	}

	paginationData := GetPaginationData(PaginationParams{
		TotalPages:   int64(math.Ceil(float64(countAirports) / float64(Limit))),
		TotalCount:   countAirports,
		CurrentPage:  int64(page),
		CurrentLimit: Limit,
	})

	err = t.Execute(w, struct {
		SidebarData    []SidebarItem
		PaginationData PaginationData
		Airports       []database.GetAirportsRow
	}{
		SidebarData:    GetSidebarData("/airports"),
		PaginationData: paginationData,
		Airports:       airports,
	})
	if err != nil {
		serverError(w, "render template", err)
		return
	}
}
