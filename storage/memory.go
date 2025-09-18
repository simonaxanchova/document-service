package storage

import (
	"document-service/models"
	"sync" // imported to use a mutex for safe concurrent access
)

// Lock/Unlock for writing
// RLock/RUnlock for reading

type MemoryStore struct {
	mu        sync.RWMutex //read-write mutex
	documents map[string]models.Document
}

// NewMemoryStore Constructor - Creates a new MemoryStore with an empty documents map
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		documents: make(map[string]models.Document),
	}
}

func (s *MemoryStore) Create(doc models.Document) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.documents[doc.ID] = doc
}

func (s *MemoryStore) GetAll() []models.Document {
	s.mu.RLock()
	defer s.mu.RUnlock()
	documents := []models.Document{}
	for _, doc := range s.documents {
		documents = append(documents, doc)
	}
	return documents
}

func (s *MemoryStore) GetByID(id string) (models.Document, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	doc, ok := s.documents[id]
	return doc, ok
}

func (s *MemoryStore) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.documents, id)
}
