package types

import "time"

type Storage struct {
	ResourceRate *Rate
	// Current storage container
	ResourceContainer *Material
	// Limit for the resource container
	ResourceContainerLimit *Material
	TimeLastCaptured       time.Time
}

func GenerateEmptyStorage() *Storage {
	return &Storage{
		ResourceContainer:      &Material{},
		ResourceRate:           &Rate{},
		ResourceContainerLimit: &Material{},
		TimeLastCaptured:       time.Now(),
	}
}

func (r *Storage) calculateOnResourceRate(timeOnCalculation *time.Time) {
	durationSinceLastCalculation := timeOnCalculation.Sub(r.TimeLastCaptured)

	resourceDiff := r.ResourceRate.getMaterialDifferenceFromDuration(&durationSinceLastCalculation)
	r.ResourceContainer.addMaterial(resourceDiff)
}

func (r *Storage) TransferResourceFrom(ext *Storage) {
	timeOnCalculation := time.Now()
	r.calculateOnResourceRate(&timeOnCalculation)
	r.TimeLastCaptured = timeOnCalculation

	r.ResourceContainer.addMaterial(ext.ResourceContainer)
}
