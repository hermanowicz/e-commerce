package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	port int = 12345
)

func main() {
	fmt.Println("Hello, E-commerce admin starting app at port:", port)

	r := chi.NewRouter()
	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World"))
	})

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      r,
	}

	// Todo: add tls support later for testing using self-sign cert
	// 	     re-use ibm tuts
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("startup failed with error:", err.Error())
	}
}
