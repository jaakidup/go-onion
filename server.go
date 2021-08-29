package main

import (
	"fmt"
	"html"
	"net/http"
	"sync"

	"github.com/jaakit/go-onion/service"
)

type Server struct {
	Name string
	Host string
	Port int
}

// TODO: Move to a seperate Kitty Service
type Kitty string

func (kitty *Kitty) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	output := []byte("Hello Kitty")

	w.Write(output)
}

// Move to separate counter service
type countHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func (s *Server) start(message string) {
	fmt.Println(message)

	http.Handle("/users", service.NewUserService())

	http.Handle("/kitty", new(Kitty))

	http.Handle("/count", new(countHandler))

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(s.Host+":"+fmt.Sprint(s.Port), nil)
}
