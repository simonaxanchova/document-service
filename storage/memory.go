package storage

import (
	"document-service/models"
	"strings"
	"sync"
)

type MemoryStore struct {
	mu        sync.RWMutex //read-write mutex
	documents map[string]models.Document
}

// NewMemoryStore Constructor
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
	text = strings.ToLower(text)
	result := false
	operator := "OR" // Default operator if not specified
	negateNext := false
	firstToken := true

	for _, token := range tokens {
		switch token {
		case "and":
			operator = "AND"
		case "or":
			operator = "OR"
		case "not":
			negateNext = true
		default:
			found := strings.Contains(text, token)
			if negateNext {
				found = !found
				negateNext = false
			}

			if firstToken {
				result = found
				firstToken = false
			} else {
				switch operator {
				case "AND":
					result = result && found
				case "OR":
					result = result || found
				}
			}
		}
	}
	return result
}
