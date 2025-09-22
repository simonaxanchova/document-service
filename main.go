package main

import (
	"document-service/handlers"
	"document-service/storage"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemoryStore()
	handler := &handlers.DocumentHandler{Store: store}

	router := mux.NewRouter()

	router.HandleFunc("/documents", handler.GetAll).Methods("GET")
	router.HandleFunc("/document/create", handler.Create).Methods("POST")
	router.HandleFunc("/document/get/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/document/delete/{id}", handler.Delete).Methods("DELETE")
	router.HandleFunc("/document/search", handler.Search).Methods("GET")

	log.Println("Starting the document-service on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
