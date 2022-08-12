package mocks

import "kata/domain"

// Notes represents a list of users who got birthday notes
type Notes []domain.User

// DummyNoteSender is dummy func that satsfies NoteSender interface
func (n *Notes) DummyNoteSender(u domain.User) {
	*n = append(*n, u)
}

// Reminders represents a map between
// a user who got reminder note and users who have a birthday
type Reminders map[domain.User][]domain.User

// DummyReminder is dummy func that satsfies UserReminder interface
func (r Reminders) DummyReminder(user domain.User, users []domain.User) {
	r[user] = users
}
