package main

import (
	"fmt"
	"time"

	"github.com/kevinliyanto/resource-simulator/types"
)

type StorageState struct {
	iron   float64
	copper float64
	coal   float64
	water  float64
	// Time in unix ms
	time time.Time
}

func getStorageState(r *types.Storage) *StorageState {
	return &StorageState{
		iron:   r.ResourceContainer.Iron,
		copper: r.ResourceContainer.Copper,
		coal:   r.ResourceContainer.Coal,
		water:  r.ResourceContainer.Water,
		time:   time.UnixMilli(r.TimeLastCaptured.UnixMilli()),
	}
}

func (s *StorageState) printStatus() string {
	return fmt.Sprintf("%+v at timestamp %v", s, s.time.UnixMilli())
}

type StorageStateRate struct {
	rate *types.Rate
	// Time diff in unix seconds
	time time.Duration
}

func (final *StorageState) deltaStorageState(initial *StorageState, offset *types.Material) *StorageStateRate {
	timeDiff := final.time.Sub(initial.time)
	timeDiffSeconds := float64(timeDiff.Seconds())

	return &StorageStateRate{
		rate: &types.Rate{
			Iron:   (final.iron - (initial.iron + offset.Iron)) / timeDiffSeconds,
			Copper: (final.copper - (initial.copper + offset.Copper)) / timeDiffSeconds,
			Coal:   (final.coal - (initial.coal + offset.Coal)) / timeDiffSeconds,
			Water:  (final.water - (initial.water + offset.Water)) / timeDiffSeconds,
		},
		time: timeDiff,
	}
}

func (s *StorageStateRate) printStatus() string {
	return fmt.Sprintf("%+v within period of %v seconds", s.rate, s.time.Seconds())
}

func (finalState *StorageStateRate) getRateDrift(originalRate *types.Rate) *types.Rate {
	return &types.Rate{
		Iron:   finalState.rate.Iron - originalRate.Iron,
		Copper: finalState.rate.Copper - originalRate.Copper,
		Coal:   finalState.rate.Coal - originalRate.Coal,
		Water:  finalState.rate.Water - originalRate.Water,
	}
}
