package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Config struct {
	UnspotURL   string
	UnspotToken string
	SpotID      string
}

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

func Handler(ctx context.Context) (*Response, error) {
	var cfg Config
	var ok bool

	cfg.UnspotURL, ok = os.LookupEnv("UNSPOT_URL")
	if !ok {
		return &Response{StatusCode: 500, Body: "UNSPOT_URL is empty"}, nil
	}

	cfg.UnspotToken, ok = os.LookupEnv("UNSPOT_TOKEN")
	if !ok {
		return &Response{StatusCode: 500, Body: "UNSPOT_TOKEN is empty"}, nil
	}

	cfg.SpotID, ok = os.LookupEnv("SPOT_ID")
	if !ok {
		return &Response{StatusCode: 500, Body: "SPOT_ID is empty"}, nil
	}

	nextDate, err := GetNextDate()
	if err != nil {
		return nil, err
	}

	err = BookSeat(cfg, *nextDate)
	if err != nil {
		return nil, err
	}

	return &Response{StatusCode: 200, Body: fmt.Sprintf("seat booked at %s", nextDate)}, nil
}
