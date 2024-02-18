package types

type Material struct {
	Iron   float64
	Copper float64
	Coal   float64
	Water  float64
}

func (r *Material) addMaterial(m *Material) {
	r.Iron += m.Iron
	r.Copper += m.Copper
	r.Coal += m.Coal
	r.Water += m.Water
}
