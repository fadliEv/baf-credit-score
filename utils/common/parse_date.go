package common

import (
	"fmt"
	"time"
)

var layoutDate = "02-01-2006"

func ParseDate(dateStr string) (time.Time,error) {	
	parseDate, err := time.Parse(layoutDate,dateStr)
	if err != nil {
		return time.Time{},fmt.Errorf("failed to parse date: %w",err)
	}
	return parseDate,nil
}

func FormatDateString(date time.Time) string {	
	return date.Format(layoutDate)
}
