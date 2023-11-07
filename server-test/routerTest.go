package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format("2006-01-02 15:00:00"), r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/user/"):]
	tmpl, err := template.New("/user/").Parse("<html><body><h1>{{.}}</h1></body></html>")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello , world")
	})

	mux.HandleFunc("/user/", greetHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggingMiddleWare(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
