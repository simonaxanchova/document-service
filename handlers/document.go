package handlers

import (
	"document-service/models"
	"document-service/storage"
	"encoding/json"
	"net/http"
)

type DocumentHandler struct {
	Store *storage.MemoryStore
}

func (h *DocumentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var doc models.Document
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.Store.Create(doc)
	w.Header().Set("Content-Type", "application/json")
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

func (h *DocumentHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}
	results := h.Store.Search(query)
	json.NewEncoder(w).Encode(results)
}
