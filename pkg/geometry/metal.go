package geometry

// Metal defines the metal material
type Metal struct {
	A    Vector
	Fuzz float64
}

// Scatter implementation
func (m Metal) Scatter(input Ray, hit HitRecord) (bool, Ray) {
	direction := input.Direction.Reflect(hit.Normal)
	bouncedRay := Ray{hit.Point, direction.Add(RandomInUnitSphere().Scale(m.Fuzz))}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}

// Albedo implementation
func (m Metal) Albedo() Vector {
	return m.A
}
