package main

import (
	"fmt"
	"time"
)

func main() {
	date := "10.02.2004"
	givenDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is %s\n", givenDate.Format("02-01-2006"), FindWeekday(givenDate))
}

func FindWeekday(date time.Time) (weekday string) {
	weekday = date.Weekday().String()
	return

}