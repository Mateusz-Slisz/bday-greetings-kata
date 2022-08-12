package domain_test

import (
	"fmt"
	"kata/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testsHasBday = []struct {
	day         time.Time
	dateOfBirth time.Time
	expected    bool
}{
	{
		day:         time.Date(2022, 2, 28, 0, 0, 0, 0, time.UTC),
		dateOfBirth: time.Date(1996, 2, 29, 0, 0, 0, 0, time.UTC),
		expected:    true,
	},
	// user born on YYYY/02/29 has bday on YYYY/02/28
	{
		day:         time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
		dateOfBirth: time.Date(1996, 2, 29, 0, 0, 0, 0, time.UTC),
		expected:    false,
	},
	{
		day:         time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC),
		dateOfBirth: time.Date(2000, 1, 15, 0, 0, 0, 0, time.UTC),
		expected:    true,
	},
	{
		day:         time.Date(2022, 1, 16, 0, 0, 0, 0, time.UTC),
		dateOfBirth: time.Date(2000, 1, 15, 0, 0, 0, 0, time.UTC),
		expected:    false,
	},
}

func TestUserHasBday(t *testing.T) {
	for _, testCase := range testsHasBday {
		// GIVEN
		user := domain.User{DateOfBirth: testCase.dateOfBirth}
		// WHEN
		hasBday := user.HasBday(testCase.day)
		// THEN
		assert.Equal(t, hasBday, testCase.expected)
	}
}

func TestUserFullName(t *testing.T) {
	// GIVEN
	user := domain.User{FirstName: "Matthew", LastName: "McConaughey"}
	fmt.Println(user.HasBday(time.Now()))
	// WHEN & THEN
	assert.Equal(t, user.FullName(), "Matthew McConaughey")
}
