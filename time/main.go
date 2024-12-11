package main

import (
	"fmt"
	"time"
)

func main() {

	withinTimeRange := isWithinTimeRange(5*time.Minute, time.Now())

	fmt.Println("withinTimeRange", withinTimeRange)

	timestamp := "2023-02-06 23:49:01"
	ts, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ts)

	withinTimeRange2 := isWithinTimeRange(5*time.Minute, ts)
	fmt.Println("withinTimeRange2", withinTimeRange2)

	minusMonth := time.Now().AddDate(0, -1, 0)

	fmt.Println("minusMonth:", minusMonth)
	withinTimeRange3 := isWithinTimeRange(5*time.Minute, minusMonth)
	fmt.Println("withinTimeRange3", withinTimeRange3)

	minus4Minutes := time.Now().Add(-4 * time.Minute)

	fmt.Println("minus4Minutes:", minus4Minutes)
	withinTimeRange4 := isWithinTimeRange(5*time.Minute, minus4Minutes)
	fmt.Println("withinTimeRange4", withinTimeRange4)

}

func isWithinTimeRange(timeRange time.Duration, t time.Time) bool {
	now := time.Now()
	diff := now.Sub(t)

	// Check if the time is within the range and in the past
	return diff >= 0 && diff <= timeRange
}
