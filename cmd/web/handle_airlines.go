package main

import (
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleAirlines(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/airlines.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	countAirlines, err := s.db.Queries.CountAirlines(r.Context())
	if err != nil {
		serverError(w, "count airlines", err)
		return
	}

	airlines, err := s.db.Queries.GetAirlines(r.Context(), database.GetAirlinesParams{
		Limit:  Limit,
		Offset: Limit * int64(page),
	})
	if err != nil {
		serverError(w, "get airlines", err)
		return
	}

	paginationData := GetPaginationData(PaginationParams{
		TotalPages:   int64(math.Ceil(float64(countAirlines) / float64(Limit))),
		TotalCount:   countAirlines,
		CurrentPage:  int64(page),
		CurrentLimit: Limit,
	})

	err = t.Execute(w, struct {
		SidebarData    []SidebarItem
		PaginationData PaginationData
		Airlines       []database.Airline
	}{
		SidebarData:    GetSidebarData("/airlines"),
		PaginationData: paginationData,
		Airlines:       airlines,
	})
	if err != nil {
		serverError(w, "render template", err)
		return
	}
}
