package main

import (
	"fmt"
	"net/http"
)

const port = ":8007"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

// middleware for http.HandleFunc
func logPath(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request from: %s\n", r.URL.Path)
		f(w, r)
	}
}

// middleware for http.Handle
func logAndServe(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request from: %s\n", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/hello", logPath(hello))

	staticHandler := http.FileServer(http.Dir("."))

	http.Handle("/", staticHandler)
	// http.Handle("/", logAndServe(staticHandler))

	fmt.Printf("Listening on http://localhost%s \n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("error: ", err)
	}
}
