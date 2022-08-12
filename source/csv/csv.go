package csv

import (
	"errors"
	"kata/domain"
)

const (
	fieldsNumber      = 4
	defaultDateFormat = "2006/01/02"
)

var (
	errInvalidFields = errors.New("data has invalid fields")
	errMissingData   = errors.New("file doesn't have data")
)

// UserSource represents source of users from csv file
type UserSource struct {
	Filename string
}

// GetUsers returns users from csv file
func (us UserSource) GetUsers() ([]domain.User, error) {
	data, err := readCSV(us.Filename)
	if err != nil {
		return nil, err
	}

	return dataToUser(data)
}
