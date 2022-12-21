package main

import (
	"backend/adapter/zinc"
	"backend/shared/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		//AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		//AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Origin", "User-Agen"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Compress(5))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/healt", func(r chi.Router) {
		r.Get("/", healt)
	})
	r.Route("/mail", func(r chi.Router) {
		r.Get("/search", searchContent)
		r.Post("/search/id", searchMail)
	})
	port := utils.CheckParam()
	log.Printf("Listen on %v", port)
	http.ListenAndServe(":"+port, r)
}

func healt(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, "Its ok !!")
}

// Search Emails
func searchContent(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("q")
	page, _ := strconv.Atoi(query.Get("page"))
	maxResults, _ := strconv.Atoi(query.Get("max_results"))
	//result := zinc.SearchByContentEs(term)
	result := zinc.SearchByContentAPI(term, page, maxResults)
	response := zinc.CreateResponseSearchService(result)
	render.JSON(w, r, response)
}

// Search Email
func searchMail(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	result := zinc.SearchByIdAPI(id)
	response := zinc.CreateResponseSearchIdService(result)
	render.JSON(w, r, response)
}
