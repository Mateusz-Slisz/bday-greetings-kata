package main

import (
	"kata/notifications/stdout"
	"kata/services"
	"kata/source/csv"
	"time"
)

func main() {
	csvSource := csv.UserSource{Filename: "input.csv"}
	tday := time.Now()

	err := services.BirthdayHandler(
		csvSource,
		stdout.SendNote,
		stdout.RemindUser,
		tday,
	)
	if err != nil {
		panic(err)
	}
}
