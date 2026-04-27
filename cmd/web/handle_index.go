package main

import (
	"net/http"
	"text/template"

	"github.com/anibaka9/faresniper/internal/database"
)

func (s Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./cmd/web/html/index.html", "./cmd/web/html/sidebar.html")
	if err != nil {
		serverError(w, "parse template", err)
		return
	}

	stats, err := s.db.Queries.GetStats(r.Context())
	if err != nil {
		serverError(w, "get stats", err)
		return
	}

	err = t.Execute(w, struct {
		SidebarData []SidebarItem
		Stats       database.GetStatsRow
	}{
		SidebarData: GetSidebarData("/"),
		Stats:       stats,
	})
	if err != nil {
		serverError(w, "render template", err)
		return
	}
}
