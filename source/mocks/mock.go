package mocks

import "kata/domain"

// DummyUserSource represents list of users
type DummyUserSource []domain.User

// GetUsers is dummy func that satsfies UserSource interface
func (s DummyUserSource) GetUsers() ([]domain.User, error) {
	return s, nil
}
