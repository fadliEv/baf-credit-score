package utils

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (time.Time,error) {
	layout := "02-01-2006"
	parseDate, err := time.Parse(layout,dateStr)
	if err != nil {
		return time.Time{},fmt.Errorf("failed to parse date: %w",err)
	}
	return parseDate,nil
}