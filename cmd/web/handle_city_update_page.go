package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleCityUpdatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/city_update.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cityId, err := strconv.Atoi(chi.URLParam(r, "city_id"))
	if err != nil {
		http.Error(w, "wrong city id", http.StatusBadRequest)
		return
	}

	city, err := s.db.Queries.GetCity(r.Context(), int64(cityId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	countries, err := s.db.Queries.GetAllCountries(r.Context())

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
