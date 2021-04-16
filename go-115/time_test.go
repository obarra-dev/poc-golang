package main_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
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

func TestTimeDate(t *testing.T) {
	dateTime := time.Date(2021, 4, 4, 0, 0, 0, 0, time.UTC)
	if dateTime.Year() != 2021 || dateTime.Month() != 4 || dateTime.Day() != 4 {
		t.Error("Error ", dateTime)
	}
}

func TestTimeParce(t *testing.T) {
	dateTime, err := time.Parse(time.RFC3339, "2020-11-01T14:00:00.000Z")
	if err != nil || dateTime.Year() != 2020 || dateTime.Month() != 11 || dateTime.Day() != 1 {
		t.Error("Error ", dateTime)
	}
}

func TestTimeFormat(t *testing.T) {
	dateTime, _ := time.Parse(time.RFC3339, "2020-11-01T14:00:00.000Z")
	r := dateTime.Format(time.ANSIC)
	if r != "Sun Nov  1 14:00:00 2020" {
		t.Error("Error ", r)
	}

	r = dateTime.Format(time.UnixDate)
	if r != "Sun Nov  1 14:00:00 UTC 2020" {
		t.Error("Error ", r)
	}

	r = dateTime.Format(time.Kitchen)
	if r != "2:00PM" {
		t.Error("Error ", r)
	}

	r = dateTime.Format("MST Jan 2 MON 2006 15:04:05")
	if r != "UTC Nov 1 MON 2020 14:00:00" {
		t.Error("Error ", r)
	}
}

// RFC3339 is ISO8601
func TestTimeISO8601(t *testing.T) {
	dateTime, _ := time.Parse(time.RFC3339, "2020-11-03T14:00:00.000Z")
	dateTimeString := dateTime.Format(time.RFC3339)
	if dateTimeString != "2020-11-03T14:00:00Z" {
		t.Error("test Error", dateTimeString)
	}

	const laytoutISO = "2006-01-02"
	dateTimeString = dateTime.Format(laytoutISO)
	if dateTimeString != "2020-11-03" {
		t.Error("test Error", dateTimeString)
	}
}

//Second is type Duration, Duration is int64
func TestTimeDuration(t *testing.T) {
	var oneSecond time.Duration = time.Second
	if oneSecond != 1000000000 || oneSecond.String() != "1s" {
		t.Errorf("Error %d, %s", oneSecond, oneSecond.String())
	}
}

func TestTimeAddDate(t *testing.T) {
	start := time.Date(2020, 4, 4, 0, 0, 0, 0, time.UTC)
	after := start.AddDate(0, 2, 0)
	r := time.Date(2020, 6, 4, 0, 0, 0, 0, time.UTC)
	if after != r {
		t.Error("Error ", after, r)
	}

	before := start.AddDate(0, -2, 0)
	r = time.Date(2020, 2, 4, 0, 0, 0, 0, time.UTC)
	if before != r {
		t.Error("Error ", before, r)
	}
}

func TestTimeDurationSince(t *testing.T) {
	now := time.Now()
	before := now.AddDate(0, 0, -1)

	var elapsed time.Duration = time.Since(before)
	r := fmt.Sprintf("%.f", elapsed.Hours())
	if r != "24" {
		t.Error("Error ", r)
	}
}

func TestTimeDurationUntil(t *testing.T) {
	now := time.Now()
	future := now.AddDate(0, 0, 1)

	var elapsed time.Duration = time.Until(future)
	r := fmt.Sprintf("%.f", elapsed.Hours())
	if r != "24" {
		t.Error("Error ", r)
	}
}
