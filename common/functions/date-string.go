package functions

import (
	"fmt"
	"time"
)

func StringToDate(date string) (time.Time, error) {
	format := "02/01/2006 15:04:05"
	return time.Parse(format, date)
}

func DateToString() string {
	t := time.Now()
	return fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
}
