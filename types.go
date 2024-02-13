package main

import (
	math_rand "math/rand"
)

type Material struct {
	Iron   float64
	Copper float64
	Coal   float64
	Water  float64
}

type Resource struct {
	resource      *Material
	resourceLimit *Material
}

func GenerateRandomResource() *Resource {
	defaultLimit := Material{
		Iron:   80000.0,
		Copper: 80000.0,
		Coal:   80000.0,
		Water:  24000.0,
	}

	return &Resource{
		resource: &Material{
			Iron:   defaultLimit.Iron * math_rand.Float64(),
			Copper: defaultLimit.Copper * math_rand.Float64(),
			Coal:   defaultLimit.Coal * math_rand.Float64(),
			Water:  defaultLimit.Water * math_rand.Float64(),
		},
		resourceLimit: &defaultLimit,
	}
}

func (r *Resource) AddMaterial(m *Material) {
	r.resource.Iron += m.Iron
	r.resource.Copper += m.Copper
	r.resource.Coal += m.Coal
	r.resource.Water += m.Water
}
