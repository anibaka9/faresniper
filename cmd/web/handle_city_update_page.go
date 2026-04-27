package main

import (
	"net/http"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleCityUpdatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/city_update.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	iata := chi.URLParam(r, "city_iata")
	if iata == "" {
		http.Error(w, "wrong city id", http.StatusBadRequest)
		return
	}

	city, err := s.db.Queries.GetCity(r.Context(), iata)
	if err != nil {
		serverError(w, "get city "+iata, err)
		return
	}

	countries, err := s.db.Queries.GetAllCountries(r.Context())
	if err != nil {
		serverError(w, "get countries", err)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		City        database.GetCityRow
		Countries   []database.GetAllCountriesRow
	}{
		SidebarData: GetSidebarData("/cities"),
		City:        city,
		Countries:   countries,
	})
	if err != nil {
		serverError(w, "render template", err)
		return
	}
}
