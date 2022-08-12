package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func readCSV(filename string) ([][]string, error) {
	f, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return [][]string{}, err
	}
	defer check(f.Close)

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func check(f func() error) {
	if err := f(); err != nil {
		fmt.Println("Received error:", err)
	}
}
