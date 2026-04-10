package main

import (
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleCountries(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/countries.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 0
	}

	countCountries, err := s.db.Queries.CountCountries(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	countries, err := s.db.Queries.GetCountries(r.Context(), database.GetCountriesParams{
		Limit:  Limit,
		Offset: Limit * int64(page),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	paginationData := GetPaginationData(PaginationParams{
		TotalPages:   int64(math.Ceil(float64(countCountries) / float64(Limit))),
		TotalCount:   countCountries,
		CurrentPage:  int64(page),
		CurrentLimit: Limit,
	})

	err = t.Execute(w, struct {
		SidebarData    []SidebarItem
		PaginationData PaginationData
		Countries      []database.Country
	}{
		SidebarData:    GetSidebarData("/countries"),
		PaginationData: paginationData,
		Countries:      countries,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
