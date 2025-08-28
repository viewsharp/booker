package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const IsDayOffUrl = "https://isdayoff.ru"

func GetNextDate() (*time.Time, error) {
	now := time.Now()
	targetDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if now.Hour() >= 12 {
		targetDate = targetDate.AddDate(0, 0, 1)
	}

	for {
		targetDate = targetDate.AddDate(0, 0, 1)

		isDayOff, err := IsDayOff(targetDate)
		if err != nil {
			return nil, fmt.Errorf("check is day off: %w", err)
		}

		if !isDayOff {
			return &targetDate, nil
		}
	}
}

func IsDayOff(date time.Time) (bool, error) {
	url := fmt.Sprintf("%s/%s", IsDayOffUrl, date.Format("20060102"))

	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("get response from %s: %v", IsDayOffUrl, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("parse response from %s: %v", IsDayOffUrl, err)
	}

	return string(body) == "1", nil
}
