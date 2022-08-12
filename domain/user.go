package domain

import (
	"strings"
	"time"
)

// User represents user data
type User struct {
	FirstName   string
	LastName    string
	Email       string
	DateOfBirth time.Time
}

// FullName returns full name of user
func (u User) FullName() string {
	return strings.Join([]string{u.FirstName, u.LastName}, " ")
}

// HasBday returns boolean when user has bday in day
// special case: user born on YYYY/02/29 has bday on YYYY/02/28
func (u User) HasBday(day time.Time) bool {
	dob := u.DateOfBirth
	if dob.Month() == 2 && dob.Day() == 29 {
		dob = dob.AddDate(0, 0, -1)
	}

	return dob.Month() == day.Month() && dob.Day() == day.Day()
}
