// Package repository = data access layer (C# ContactStore / future DB).
// Controllers/handlers never talk to storage directly — they go through services → repo.
package repository

import (
	"sync"
	"time"

	"github.com/adityax28/fastport/internal/models"
)

// ContactRepository = in-memory "database" for contact messages.
// Later you can swap this for Postgres/SQLite without changing handlers.
type ContactRepository struct {
	mu       sync.Mutex
	messages []models.StoredContact
}

// NewContactRepository constructs an empty store (DI-style factory).
func NewContactRepository() *ContactRepository {
	return &ContactRepository{
		messages: make([]models.StoredContact, 0),
	}
}

// Add persists one contact message (thread-safe).
func (r *ContactRepository) Add(req models.ContactRequest) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.messages = append(r.messages, models.StoredContact{
		At:      time.Now().UTC(),
		Request: req,
	})
}

// Count returns how many messages are stored.
func (r *ContactRepository) Count() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.messages)
}
