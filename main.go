package main

import (
	"document-service/handlers"
	"document-service/storage"
	"log"
	"net/http"
)

func main() {
	store := storage.NewMemoryStore()
	handler := &handlers.DocumentHandler{Store: store}

	http.HandleFunc("/documents", handler.GetAll)
	http.HandleFunc("/document/create", handler.Create)
	http.HandleFunc("/document/get", handler.GetByID)
	http.HandleFunc("/document/delete", handler.Delete)
	http.HandleFunc("/document/search", handler.Search)

	log.Println("Starting the document-service on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
