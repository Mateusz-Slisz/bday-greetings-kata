package services

import (
	"kata/domain"
	"time"
)

func getBdayUsers(users []domain.User, day time.Time) []domain.User {
	bdayUsers := []domain.User{}
	for _, user := range users {
		if user.HasBday(day) {
			bdayUsers = append(bdayUsers, user)
		}
	}
	return bdayUsers
}

func getExcludedUsers(bdayUsers []domain.User, u domain.User) []domain.User {
	excluded := []domain.User{}
	for _, bdayUser := range bdayUsers {
		if bdayUser != u {
			excluded = append(excluded, bdayUser)
		}
	}
	return excluded
}
