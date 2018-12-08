package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vjeantet/jodaTime"
)

func Test_time_strings(t *testing.T) {
	date := jodaTime.Format("YYYY.MM.dd", time.Now())
	fmt.Println(date)

	date = jodaTime.Format("YYYY-MM-dd hh:m:ss", time.Now())
	fmt.Println(date)

	dateTime, _ := jodaTime.Parse("dd/MMMM/yyyy:HH:mm:ss", "30/August/2015:21:44:25")
	fmt.Println(dateTime.String())

}
