package csv

import (
	"kata/domain"
	"strings"
	"time"
)

func dataToUser(data [][]string) ([]domain.User, error) {
	if len(data) < 1 {
		return nil, errMissingData
	}

	var users []domain.User
	for _, line := range data[1:] {
		user, err := lineToUser(line)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func lineToUser(line []string) (domain.User, error) {
	if len(line) != fieldsNumber {
		return domain.User{}, errInvalidFields
	}

	dof, err := time.Parse(defaultDateFormat, strings.Trim(line[2], " "))
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		FirstName:   strings.Trim(line[1], " "),
		LastName:    strings.Trim(line[0], " "),
		DateOfBirth: dof,
		Email:       strings.Trim(line[3], " "),
	}, nil
}
