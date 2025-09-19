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

func TestGetNonExistent(t *testing.T) {
	store := storage.NewMemoryStore()

	_, ok := store.GetByID("999")
	if ok {
		t.Errorf("Expected false, got true for non existing document")
	}
}

func TestCreateMultiple(t *testing.T) {
	store := storage.NewMemoryStore()

	docs := []models.Document{
		{ID: "1", Name: "Test 1", Description: "I am testing"},
		{ID: "2", Name: "Test 2", Description: "I am testing again"},
	}

	for _, doc := range docs {
		store.Create(doc)
	}

	for _, doc := range docs {
		got, ok := store.GetByID(doc.ID)
		if !ok || got.Name != doc.Name {
			t.Errorf("Expected document '%s', got '%v'", doc.Name, got.Name)
		}
	}
}

func TestUpdateDocument(t *testing.T) {
	store := storage.NewMemoryStore()
	doc := models.Document{ID: "1", Name: "Test 1", Description: "I am testing"}
	store.Create(doc)

	doc.Name = "NewName"
	store.Create(doc)

	got, ok := store.GetByID("1")
	if !ok || got.Name != "NewName" {
		t.Errorf("Expected document 'NewName', got '%v'", got.Name)
	}
}

func TestDeleteDocument(t *testing.T) {
	store := storage.NewMemoryStore()
	doc := models.Document{ID: "1", Name: "Test"}
	store.Create(doc)

	store.Delete("1")

	_, ok := store.GetByID("1")
	if ok {
		t.Errorf("Expected document to be deleted")
	}
}
