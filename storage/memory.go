package storage

import (
	"document-service/models"
	"strings"
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

func (s *MemoryStore) Search(query string) []models.Document {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tokens := strings.Fields(strings.ToLower(query))

	results := []models.Document{}
	for _, doc := range s.documents {
		text := strings.ToLower(doc.Name + " " + doc.Description)

		if matchesQuery(text, tokens) {
			results = append(results, doc)
		}
	}
	return results
}

func matchesQuery(text string, tokens []string) bool {
	result := false
	operator := "OR" // Default operator if not specified

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch token {
		case "and":
			operator = "AND"
		case "or":
			operator = "OR"
		case "not":
			operator = "NOT"
		default:
			found := strings.Contains(text, token)

			switch operator {
			case "AND":
				result = result && found
			case "OR":
				result = result || found
			case "NOT":

				result = result && !found
			default:
				result = found
			}
		}
	}
	return result
}
