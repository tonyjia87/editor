package main

import (
	"github.com/gorilla/handlers"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/shurcooL/httpgzip"
	"github.com/tonyjia87/editor/frontend"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type Config struct {
	RelativeRoot string
	BindAddr     string
}

var config = &Config{
	RelativeRoot: "/",
	BindAddr:     ":8080",
}

func noCacheControl(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tf := template.New("index.html").Delims("[[", "]]")
	t := template.Must(vfstemplate.ParseFiles(frontend.Assets,
		tf,
		"/templates/index.html"))

	t.Execute(w, config)

	//File := template.New(vfstemplate.ParseFiles(frontend.Assets, nil, "/templates/index.html")).Delims("[[", "]]").
}

func setupRoutes(relativeroot string) *http.ServeMux {
	router := http.NewServeMux()

	staticHandler := noCacheControl(httpgzip.FileServer(frontend.Assets, httpgzip.FileServerOptions{IndexHTML: true}))

	staticHandler = http.StripPrefix(relativeroot+"vfs/", staticHandler)
	router.Handle(relativeroot+"vfs/", staticHandler)

	router.HandleFunc(relativeroot+"", indexHandler)

	return router
}

func setupServer(config *Config, addr string, logger *log.Logger) *http.Server {
	router := setupRoutes(config.RelativeRoot)
	loggingRouter := handlers.LoggingHandler(os.Stderr, router)

	server := http.Server{
		Addr:         addr,
		Handler:      loggingRouter,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server
}

func startServer(config *Config, bindAddr string) {
	loggerHTML := log.New(os.Stdout, "", log.LstdFlags)
	loggerHTML.Printf("Server start, relative-root: %s, bind-addr: %s\n", config.RelativeRoot, bindAddr)

	server := setupServer(config, bindAddr, loggerHTML)

	server.ListenAndServe()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go startServer(config, config.BindAddr)
	wg.Wait()
}
