package stdout

import (
	"fmt"
	"kata/domain"
	"strings"
)

// RemindUser reminds user about birthdays through stdout
func RemindUser(u domain.User, users []domain.User) {
	userNames := getUserNames(users)
	fmt.Printf(
		"Subject: Birthday Reminder!\n"+
			"Dear %s,\n"+
			"Today is %s's birthday\n"+
			"Don't forget to send him a message!\n\n",
		u.FirstName,
		strings.Join(userNames, ", "),
	)
}

func getUserNames(users []domain.User) []string {
	userNames := []string{}
	for _, user := range users {
		userNames = append(userNames, user.FullName())
	}
	return userNames
}
