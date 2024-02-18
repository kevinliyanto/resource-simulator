package types

import (
	"fmt"
	"time"
)

// Rate per 1000 ms or time.ParseDuration("1s")
type Rate struct {
	Iron   float64
	Copper float64
	Coal   float64
	Water  float64
}

func (resourceRate *Rate) getMaterialDifferenceFromDuration(d *time.Duration) *Material {
	// Calculated up to nanosecond for most precise calculation
	duration := float64(d.Nanoseconds()) / 1e9

	return &Material{
		Iron:   resourceRate.Iron * duration,
		Copper: resourceRate.Copper * duration,
		Coal:   resourceRate.Coal * duration,
		Water:  resourceRate.Water * duration,
	}
}

func (rate *Rate) PrintStatus() string {
	return fmt.Sprintf("%+v", rate)
}
