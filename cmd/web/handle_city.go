package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
	"github.com/go-chi/chi/v5"
)

func (s Server) handleCity(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/city.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	iata := chi.URLParam(r, "city_iata")
	fmt.Println(iata)
	if iata == "" {
		http.Error(w, "wrong city iata", http.StatusBadRequest)
		return
	}

	city, err := s.db.Queries.GetCity(r.Context(), iata)
	if err != nil {
		serverError(w, "get city "+iata, err)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		City        database.GetCityRow
	}{
		SidebarData: GetSidebarData("/cities"),
		City:        city,
	})
	if err != nil {
		serverError(w, "render template", err)
		return
	}
}
