package geometry

// Ray is a struct defining a ray consisting of the ray origin and the ray direction
type Ray struct {
	Origin    Vector
	Direction Vector
}

// PointOnRay provides the computation of a point along the ray
func (r Ray) PointOnRay(t float64) Vector {
	return r.Origin.Add(r.Direction.Scale(t))
}

// TODO color nad hitSphere honestly probably need to be extracted out into there own thing
// Color takes a ray value receiver and applies color to the ray
func (r Ray) Color() Vector {
	if r.hitSphere(Vector{X: 0, Y: 0, Z: -1}, 0.5) {
		return Vector{X: 1, Y: 0, Z: 0}
	}
	unitDirection := r.Direction.Normalize()
	t := 0.5 * (unitDirection.Y + 1.0)

	scaledVector1 := Vector{X: 1.0, Y: 1.0, Z: 1.0}.Scale(1.0 - t)
	scaledVector2 := Vector{X: 0.5, Y: 0.7, Z: 1.0}.Scale(t)

	finalVector := scaledVector1.Add(scaledVector2)

	return finalVector
}

func (r Ray) hitSphere(v Vector, radius float64) bool {
	oc := r.Origin.Subtract(v)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - radius * radius
	discriminant := b * b - 4 * a * c
	return discriminant > 0
}
