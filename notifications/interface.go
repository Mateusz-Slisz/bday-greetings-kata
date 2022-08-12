package notifications

import "kata/domain"

// ReminderFunc is an interface for func that will remind users about birthdays
type ReminderFunc func(domain.User, []domain.User)

// NoteSenderFunc is an interface for func that will send a Birthday note
type NoteSenderFunc func(domain.User)
