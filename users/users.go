package users

import (
	"errors"
	"fmt"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

// Store is a very simple in memory database to store our users.
// It is protects by a read-write mutex. This stops 2 goroutines
// from modifying the underlying map at the same time (lock)
// This is because maps are not safe for concurrent use in Go.
type Store struct {
	rwm *sync.RWMutex
	m   map[string]string
}

var (
	DB                   = newDB()
	ErrUserAlreadyExists = errors.New("users: Username already exists")
)

// newDB is a helper methof to init out in-memory DB when program
// starts
func newDB() *Store {
	return &Store{
		rwm: &sync.RWMutex{},
		m:   make(map[string]string),
	}
}

// NewUser creates user with hashed password
func NewUser(username, password string) error {
	err := exists(username)
	if err != nil {
		return err
	}

	err = SetPassword(username, password)
	if err != nil {
		return err
	}
	return nil
}

func SetPassword(username, password string) error {
	DB.rwm.Lock()
	defer DB.rwm.Unlock()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	DB.m[username] = string(hashedPassword)
	fmt.Println("SetPassword:", DB.m[username])
	return nil
}

func AuthenticateUser(username, password string) error {
	DB.rwm.RLock()
	defer DB.rwm.RUnlock()

	hashedPassword := DB.m[username]
	fmt.Println("pass:", password)
	fmt.Println("encp:", hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}

// exists is a helper function for checking if usernames are unique
func exists(username string) error {
	DB.rwm.RLock()
	defer DB.rwm.RUnlock()

	if DB.m[username] != "" {
		return ErrUserAlreadyExists
	}
	return nil
}
