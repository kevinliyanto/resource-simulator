package main

import (
	"fmt"

	"github.com/kevinliyanto/resource-simulator/types"
)

type StorageState struct {
	material *types.Material
	// Time in unix ms
	time int64
}

func getStorageState(r *types.Storage) *StorageState {
	return &StorageState{
		material: &types.Material{
			Iron:   r.ResourceContainer.Iron,
			Copper: r.ResourceContainer.Copper,
			Coal:   r.ResourceContainer.Coal,
			Water:  r.ResourceContainer.Water,
		},
		time: r.TimeLastCaptured.UnixMilli(),
	}
}

func (s *StorageState) printStatus() string {
	return fmt.Sprintf("%+v at timestamp %v", s.material, s.time)
}

type StorageStateRate struct {
	Iron   float64
	Copper float64
	Coal   float64
	Water  float64
	// Time diff in unix seconds
	time int64
}

func (final *StorageState) deltaStorageState(initial *StorageState, offset *types.Material) *StorageStateRate {
	time := (final.time - initial.time) / 1e3

	return &StorageStateRate{
		Iron:   (final.material.Iron - (initial.material.Iron + offset.Iron)) / float64(time),
		Copper: (final.material.Copper - (initial.material.Copper + offset.Copper)) / float64(time),
		Coal:   (final.material.Coal - (initial.material.Coal + offset.Coal)) / float64(time),
		Water:  (final.material.Water - (initial.material.Water + offset.Water)) / float64(time),
		time:   time,
	}
}

func (s *StorageStateRate) printStatus() string {
	return fmt.Sprintf("%+v within period of %v seconds", s, s.time)
}
