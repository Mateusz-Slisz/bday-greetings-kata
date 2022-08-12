package services

import (
	"kata/domain"
	"kata/notifications"
	"kata/source"
	"time"
)

// BirthdayHandler sends a birthday's note and reminds other users
func BirthdayHandler(
	source source.UserSourcer,
	sendFunc notifications.NoteSenderFunc,
	remindFunc notifications.ReminderFunc,
	dateTarget time.Time,
) error {
	users, err := source.GetUsers()
	if err != nil {
		return err
	}

	bdayUsers := getBdayUsers(users, dateTarget)

	handleNotes(sendFunc, bdayUsers)
	handleReminders(remindFunc, users, bdayUsers)

	return nil
}

func handleNotes(s notifications.NoteSenderFunc, bdayUsers []domain.User) {
	for _, bdayUser := range bdayUsers {
		s(bdayUser)
	}
}

func handleReminders(
	r notifications.ReminderFunc,
	users []domain.User,
	bdayUsers []domain.User,
) {
	if len(bdayUsers) == 0 {
		return
	}
	for _, user := range users {
		excluded := getExcludedUsers(bdayUsers, user)

		if len(excluded) > 0 {
			r(user, excluded)
		}
	}
}
