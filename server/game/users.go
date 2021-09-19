package game

import (
	"sync"
	"time"
)

type User struct {
	puzzle  *Puzzle
	created time.Time
}

func (u *User) GetPuzzle() *Puzzle {
	return u.puzzle
}

func NewUser(puzzle *Puzzle) *User {
	return &User{puzzle: puzzle, created: time.Now()}
}

type UserManager struct {
	lock  *sync.Mutex
	users map[string]*User
}

func NewUserManager() *UserManager {
	return &UserManager{lock: &sync.Mutex{}, users: make(map[string]*User)}
}

func (mgr *UserManager) Add(phoneNumber string, user *User) {
	// Add user
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	mgr.users[phoneNumber] = user
}

func (mgr *UserManager) Get(phoneNumber string) *User {
	// Get user
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	return mgr.users[phoneNumber]
}

func (mgr *UserManager) Remove(phoneNumber string) {
	// Delete user. can be used once the user solves their puzzle
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	delete(mgr.users, phoneNumber)
}
