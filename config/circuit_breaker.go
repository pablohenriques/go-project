package config

import (
	"log"
	"time"

	"github.com/sony/gobreaker"
)

func GetSetting() gobreaker.Settings {
	settings := gobreaker.Settings{
		Name:    "CircuitExample",
		Timeout: 3 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 2
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Changed State: %s -> %s", from, to)
		},
	}
	return settings
}
