package goratp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const defaultEndpoint = "https://api-ratp.pierre-grimaud.fr/v4"

// Schedule for a station on a line according to the destination.
type Schedule map[string][]string

// GetSchedules getter at a station on a line.
func GetSchedules(station, line, transport string) (Schedule, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", defaultEndpoint, "schedules", transport, line, station, "A+R")
	sc := make(Schedule)

	var result struct {
		Result struct {
			Schedules []struct {
				Message     string `json:"message"`
				Destination string `json:"destination"`
			} `json:"schedules"`
		} `json:"result"`
	}

	clt := &http.Client{
		Timeout: 10 * time.Second,
	}
	r, err := clt.Get(url)
	if err != nil {
		return Schedule{}, err
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return Schedule{}, err
	}

	for _, s := range result.Result.Schedules {
		sc[s.Destination] = append(sc[s.Destination], s.Message)
	}

	return sc, nil
}
