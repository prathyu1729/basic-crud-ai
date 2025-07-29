package store

import (
	"sync"
	"time"

	"basic-crud-ai/models"
)

// UserStore holds the in-memory storage for users
type UserStore struct {
	users  map[int]*models.User
	mutex  sync.RWMutex
	nextID int
}

// NewUserStore creates a new user store
func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}

// CreateUser adds a new user to the store
func (store *UserStore) CreateUser(name, email string) *models.User {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	now := time.Now()
	user := &models.User{
		ID:        store.nextID,
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	store.users[user.ID] = user
	store.nextID++
	return user
}

// GetUser retrieves a user by ID
func (store *UserStore) GetUser(id int) (*models.User, bool) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user, exists := store.users[id]
	return user, exists
}

// GetAllUsers returns all users
func (store *UserStore) GetAllUsers() []*models.User {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	users := make([]*models.User, 0, len(store.users))
	for _, user := range store.users {
		users = append(users, user)
	}
	return users
}

// UpdateUser updates an existing user
func (store *UserStore) UpdateUser(id int, name, email string) (*models.User, bool) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	user, exists := store.users[id]
	if !exists {
		return nil, false
	}

	user.Name = name
	user.Email = email
	user.UpdatedAt = time.Now()

	return user, true
}

// DeleteUser removes a user by ID
func (store *UserStore) DeleteUser(id int) bool {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	_, exists := store.users[id]
	if !exists {
		return false
	}

	delete(store.users, id)
	return true
}
