package main

import (
	"document-service/handlers"
	"document-service/storage"
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

	http.ListenAndServe(":8080", nil)
}
