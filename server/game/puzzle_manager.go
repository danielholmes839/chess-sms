package game

import (
	"sync"
)

type PuzzleManager struct {
	lock    *sync.Mutex
	puzzles map[string]*Puzzle // store the active puzzle for a given phonenumber
}

func NewPuzzleManager() *PuzzleManager {
	return &PuzzleManager{lock: &sync.Mutex{}, puzzles: make(map[string]*Puzzle)}
}

func (mgr *PuzzleManager) Add(phoneNumber string, puzzle *Puzzle) {
	// Add user
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	mgr.puzzles[phoneNumber] = puzzle
}

func (mgr *PuzzleManager) Get(phoneNumber string) *Puzzle {
	// Get user
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	return mgr.puzzles[phoneNumber]
}

func (mgr *PuzzleManager) Remove(phoneNumber string) {
	// Delete user. can be used once the user solves their puzzle
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	delete(mgr.puzzles, phoneNumber)
}
