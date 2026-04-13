package main

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleCountryUpdatePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/country_update.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	countryId, err := strconv.Atoi(chi.URLParam(r, "country_id"))
	if err != nil {
		http.Error(w, "wrong county id", http.StatusBadRequest)
		return
	}

	country, err := s.db.Queries.GetCountry(r.Context(), int64(countryId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		Country     database.Country
	}{
		SidebarData: GetSidebarData("/countries"),
		Country:     country,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
