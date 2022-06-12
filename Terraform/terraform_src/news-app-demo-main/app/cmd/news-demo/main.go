package main

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"
	"log"
	"math"
	"net/http"
	"net/url"
	"news-demo/internal/version"
	"strconv"
	"time"

	"news-demo/internal/config"
	"news-demo/internal/news"
)

//go:embed templates
var indexHTML embed.FS

//go:embed assets
var assets embed.FS

var tpl = template.Must(template.ParseFS(indexHTML, "templates/index.html"))

type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *news.Results
}

func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func searchHandler(newsapi news.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		results, err := newsapi.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &Search{
			Query:      searchQuery,
			NextPage:   nextPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(newsapi.GetPageSize()))),
			Results:    results,
		}

		if ok := !search.IsLastPage(); ok {
			search.NextPage++
		}

		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}


func main() {

	log.Println("Start application...")
	log.Println("Version: ", version.Version)
	log.Println("Build Time: ", version.BuildTime)
	log.Println("Commit", version.Commit)

	cfg := config.GetConfig()
	if cfg.ApiKey == "" {
		log.Fatal("Env: apiKey must be set")
	}

	myClient := &http.Client{Timeout: 60 * time.Second}
	newsapi := news.NewClient(myClient, cfg.ApiKey, cfg.NewsLang, 20)

	fs := http.FileServer(http.FS(assets))

	mux := http.NewServeMux()
	mux.Handle("/assets/", fs)
	mux.HandleFunc("/search", searchHandler(newsapi))
	mux.HandleFunc("/", indexHandler)

	server := &http.Server{
		Addr: cfg.BindAddr,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
