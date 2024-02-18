package types

type Material struct {
	Iron   Iron
	Copper Copper
	Coal   Coal
	Water  Water
}

func (r *Material) addMaterial(m *Material) {
	r.Iron += m.Iron
	r.Copper += m.Copper
	r.Coal += m.Coal
	r.Water += m.Water
}
