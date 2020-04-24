package main

import (
	"github.com/gorilla/handlers"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/shurcooL/httpgzip"
	"github.com/tonyjia87/editor/frontend"
	"github.com/tonyjia87/editor/middleware"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tf := template.New("index.html").Delims("[[", "]]")
	t := template.Must(vfstemplate.ParseFiles(frontend.Assets,
		tf,
		"/templates/index.html"))

	t.Execute(w, config)
}

func proxyHandler(w http.ResponseWriter, r *http.Request)  {
	transport :=  http.DefaultTransport
	outReq := new(http.Request)

	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	res, err := transport.RoundTrip(outReq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	// step 3
	for key, value := range res.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}

	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	res.Body.Close()
}

func setupRoutes(relativeroot string) *http.ServeMux {
	router := http.NewServeMux()

	staticHandler := middleware.Set(httpgzip.FileServer(frontend.Assets, httpgzip.FileServerOptions{IndexHTML: true}))

	staticHandler = http.StripPrefix(relativeroot+"vfs/", staticHandler)
	router.Handle(relativeroot+"vfs/", staticHandler)

	router.HandleFunc(relativeroot+"app.json", proxyHandler)
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
