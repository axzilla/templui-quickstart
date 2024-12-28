// store/users.go
package store

import (
	"fmt"
	"sync"
)

type User struct {
	Username string
	Email    string
	ID       int
}

type UserStore struct {
	sync.RWMutex
	users  []User
	lastID int
}

var Store = &UserStore{
	users: []User{
		{Username: "Alice", Email: "userone@test.com", ID: 1},
		{Username: "Bob", Email: "usertwo@test.com", ID: 2},
		{Username: "Charlie", Email: "userthree@test.com", ID: 3},
	},
	lastID: 3,
}

func (s *UserStore) List() []User {
	s.RLock()
	defer s.RUnlock()
	users := make([]User, len(s.users))
	copy(users, s.users)
	return users
}

func (s *UserStore) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	for i, user := range s.users {
		if user.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
