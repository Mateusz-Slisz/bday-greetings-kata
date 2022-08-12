package source

import "kata/domain"

// UserSourcer is an interface for source of user data
type UserSourcer interface {
	GetUsers() ([]domain.User, error)
}
