package main

import (
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/anibaka9/faresniper/internal/database"
	databaseclient "github.com/anibaka9/faresniper/internal/database_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const Limit int64 = 10

func main() {
	c, err := databaseclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./cmd/web/html/index.html", "./cmd/web/html/sidebar.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		stats, err := c.Queries.GetStats(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/routes", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("./cmd/web/html/routes.html", "./cmd/web/html/sidebar.html", "./cmd/web/html/pagination.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 0
		}

		countRoutes, err := c.Queries.CountRoutes(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		routes, err := c.Queries.GetRoutes(r.Context(), database.GetRoutesParams{
			Limit:  Limit,
			Offset: Limit * int64(page),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		paginationData := GetPaginationData(PaginationParams{
			TotalPages:   int64(math.Ceil(float64(countRoutes) / float64(Limit))),
			TotalCount:   countRoutes,
			CurrentPage:  int64(page),
			CurrentLimit: Limit,
		})

		err = t.Execute(w, struct {
			SidebarData    []SidebarItem
			PaginationData PaginationData
			Routes         []database.GetRoutesRow
		}{
			SidebarData:    GetSidebarData("/routes"),
			PaginationData: paginationData,
			Routes:         routes,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Handle("/tailwind/*", http.StripPrefix("/tailwind/", http.FileServer(http.Dir("./cmd/web/tailwind"))))
	http.ListenAndServe(":3000", r)
}
