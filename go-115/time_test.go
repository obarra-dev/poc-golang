package main_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	showTime()
}

func TestTimeISO8601(t *testing.T) {
	dateTime, _ := time.Parse(time.RFC3339, "2020-11-03T14:00:00.000Z")
	dateTimeString := dateTime.Format(time.RFC3339)
	if dateTimeString != "2020-11-03T14:00:00Z" {
		t.Error("test Error", dateTimeString)
	}
}

func showTime() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
