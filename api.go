package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Storage
}

func NewAPIServer(addr string, store Storage) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/todo", makeHTTPHandlerFunc(s.handleTodo)).Methods("GET", "POST")
	router.HandleFunc("/todo/{id}", makeHTTPHandlerFunc(s.handleSingleTodo)).Methods("GET", "DELETE", "PUT")

	log.Println("Server start on port ", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, router))
}
