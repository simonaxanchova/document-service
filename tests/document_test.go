package tests

import (
	"document-service/models"
	"document-service/storage"
	"testing"
)

func TestCreateAndGet(t *testing.T) {
	store := storage.NewMemoryStore()
	doc := models.Document{ID: "1", Name: "Test", Description: "I am testing"}
	store.Create(doc)

	got, ok := store.GetByID("1")
	if !ok || got.Name != "Test" {
		t.Errorf("Expected document 'Test', got '%v'", got.Name)
	}
}
