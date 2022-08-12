package services_test

import (
	"kata/domain"
	notifyMocks "kata/notifications/mocks"
	"kata/services"
	sourceMocks "kata/source/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	sinatra = domain.User{
		FirstName:   "Frank",
		LastName:    "Sinatra",
		DateOfBirth: time.Date(1915, 12, 12, 0, 0, 0, 0, time.UTC),
	}
	flowerBoy = domain.User{
		FirstName:   "Flower",
		LastName:    "Boy",
		DateOfBirth: time.Date(1966, 2, 14, 0, 0, 0, 0, time.UTC),
	}

	lovelas = domain.User{
		FirstName:   "Lovely",
		LastName:    "Lovelas",
		DateOfBirth: time.Date(1999, 2, 14, 0, 0, 0, 0, time.UTC),
	}
)

func TestBirthdayHandler(t *testing.T) {
	// GIVEN
	valentineDay := time.Date(2022, 2, 14, 0, 0, 0, 0, time.UTC)
	dummySource := sourceMocks.DummyUserSource{sinatra, flowerBoy, lovelas}
	notes := notifyMocks.Notes{}
	reminders := notifyMocks.Reminders{}

	// WHEN
	err := services.BirthdayHandler(
		dummySource,
		notes.DummyNoteSender,
		reminders.DummyReminder,
		valentineDay,
	)

	// THEN
	assert.Nil(t, err)
	assert.ElementsMatch(
		t,
		notes,
		[]domain.User{flowerBoy, lovelas},
	) // flowerBoy and lovelas got birthday note
	assert.ElementsMatch(
		t,
		reminders[sinatra],
		[]domain.User{flowerBoy, lovelas},
	) // sinatra got reminder about flowerBoy and lovelas birthday
	assert.ElementsMatch(
		t,
		reminders[flowerBoy],
		[]domain.User{lovelas},
	) // flowerBoy got reminder about lovelas birthday
	assert.ElementsMatch(
		t,
		reminders[lovelas],
		[]domain.User{flowerBoy},
	) // lovelas got reminder about flowerBoy birthday
}
