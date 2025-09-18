package handlers

import (
	"document-service/models"
	"document-service/storage"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type DocumentHandler struct {
	Store *storage.MemoryStore
}

func (h *DocumentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var doc models.Document
	fmt.Println("Tuka sum")
	json.NewDecoder(r.Body).Decode(&doc)
	doc.ID = uuid.New().String()
	h.Store.Create(doc)
	json.NewEncoder(w).Encode(doc)
}

func (h *DocumentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.Store.GetAll())
}

func (h *DocumentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if doc, ok := h.Store.GetByID(id); ok {
		json.NewEncoder(w).Encode(doc)
	} else {
		http.Error(w, "Document not found", http.StatusNotFound)
	}
}

func (h *DocumentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	h.Store.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
