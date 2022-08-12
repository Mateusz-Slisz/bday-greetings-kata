package stdout

import (
	"fmt"
	"kata/domain"
)

// SendNote sends user's birthday note to stdout
func SendNote(u domain.User) {
	fmt.Printf(
		"Subject: Happy birthday!\nHappy birthday, dear %s!\n\n",
		u.FirstName,
	)
}
