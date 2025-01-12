package main

import (
	"log"
	"net/http"
	"time"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func TimeUse(handler Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler(w, r)
		log.Printf("url: %s, time: %v", r.URL, time.Since(now))
	}
}

func TimeUse2(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("url: %s, time: %v", r.URL, time.Since(now))
	}

	return http.HandlerFunc(fn)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world"))
}

func HowAreYou(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("i'am fine"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", HelloWorld)
	// mux.HandleFunc("GET /hello", TimeUse(HelloWorld))
	mux.HandleFunc("GET /how", HowAreYou)

	srv := http.Server {
		Addr: ":8080",
		Handler: TimeUse2(mux),
	}

	srv.ListenAndServe()
}