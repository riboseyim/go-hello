package test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_date_strings(t *testing.T) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth, lastOfMonth := get_point_days(currentYear, currentMonth.String())

	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Chongqing")
	zero_day, _ := time.ParseInLocation(layout, "1949-10-01 00:00:00", loc)
	log.Println("zero_day:%s", zero_day)
}

func get_point_days(year int, m Month) (string, string) {

	firstOfMonth := time.Date(year, m, 1, 0, 0, 0, 0, now.Location())
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	fmt.Println("firstOfMonth:%s", firstOfMonth.Unix())
	fmt.Println("lastOfMonth:%s", lastOfMonth.Unix())

	return firstOfMonth, lastOfMonth
}
