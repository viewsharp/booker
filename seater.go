package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// TimeRange represents the booking time range
type TimeRange struct {
	StartDate int64 `json:"startDate"`
	EndDate   int64 `json:"endDate"`
}

// DeskRequest represents the desk booking request
type DeskRequest struct {
	Comment       *string     `json:"comment,omitempty"`
	Desk          string      `json:"desk"`
	Force         bool        `json:"force"`
	GuestEmail    *string     `json:"guestEmail,omitempty"`
	GuestName     *string     `json:"guestName,omitempty"`
	IsBookingByQr bool        `json:"isBookingByQr"`
	Organizer     *string     `json:"organizer,omitempty"`
	PlaceFrom     string      `json:"placeFrom"`
	Title         string      `json:"title"`
	TimeRanges    []TimeRange `json:"timeRanges"`
}

// BookSeat books a seat in Unspot
func BookSeat(config Config, date time.Time) error {
	client := &http.Client{Timeout: 10 * time.Second}

	requestBody := DeskRequest{
		Desk:      config.SpotID,
		PlaceFrom: "MAP_DESKTOP",
		Title:     "4.120 забронирован в 4 этаж",
		TimeRanges: []TimeRange{{
			StartDate: date.Add(time.Hour * 9).Unix(),
			EndDate:   date.Add(time.Hour * 19).Unix(),
		}},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("marshal request body: %w", err)
	}

	url := fmt.Sprintf("%sapi/bookings/desk/new", config.UnspotURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.UnspotToken)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request unspot: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("booking failed with status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Printf("Booking successful. Status: %d\n", resp.StatusCode)
	return nil
}
