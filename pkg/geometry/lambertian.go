package geometry

// TODO should be in its own package but couldn't figure out the circular dependacies

// Lambertian defines the lambertian material
type Lambertian struct {
	A Vector
}

// Scatter implementation
func (l Lambertian) Scatter(input Ray, hit HitRecord) (bool, Ray) {
	// Book uses target which adds Point but then on the return subtracts it...so idk
	// target := hit.Point.Add(Normal.Add(RandomInUnitSphere()))
	direction := hit.Normal.Add(RandomInUnitSphere())
	return true, Ray{Origin: hit.Point, Direction: direction}
}

// Albedo implementation
func (l Lambertian) Albedo() Vector {
	return l.A
}