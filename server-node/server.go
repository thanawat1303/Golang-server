package main

import (
	"fmt"
	"net/http"
	"time"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         "3001",
		Handler:      Authentication(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
