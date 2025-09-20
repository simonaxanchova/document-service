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

func TestGetAll(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Create(models.Document{ID: "1", Name: "Doc1", Description: ""})
	store.Create(models.Document{ID: "2", Name: "Doc2", Description: ""})

	results := store.GetAll()
	if len(results) != 2 {
		t.Errorf("Expected 2 documents, got %d", len(results))
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

func TestSearchAND(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Create(models.Document{ID: "1", Name: "Apple Orange", Description: ""})
	store.Create(models.Document{ID: "2", Name: "Apple", Description: ""})
	store.Create(models.Document{ID: "3", Name: "Orange", Description: ""})

	results := store.Search("apple AND orange")
	if len(results) != 1 || results[0].ID != "1" {
		t.Errorf("Expected only document 1, got %v", results)
	}
}

func TestSearchOR(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Create(models.Document{ID: "1", Name: "Apple", Description: ""})
	store.Create(models.Document{ID: "2", Name: "Orange", Description: ""})

	results := store.Search("apple OR orange")
	if len(results) != 2 {
		t.Errorf("Expected 2 documents, got %d", len(results))
	}
}

func TestSearchNOT(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Create(models.Document{ID: "1", Name: "Apple", Description: ""})
	store.Create(models.Document{ID: "2", Name: "Banana", Description: ""})

	results := store.Search("apple NOT banana")
	if len(results) != 1 || results[0].ID != "1" {
		t.Errorf("Expected only document 1, got %v", results)
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	store := storage.NewMemoryStore()
	store.Create(models.Document{ID: "1", Name: "Apple", Description: ""})

	results := store.Search("APPLE")
	if len(results) != 1 {
		t.Errorf("Expected 1 document, got %d", len(results))
	}
}
