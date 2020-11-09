package geometry

// Ray is a struct defining a ray consisting of the ray origin and the ray direction
type Ray struct {
	Origin    Vector
	Direction Vector
}

// PointOnRay provides the computation of a point along the ray
func (r *Ray) PointOnRay(t float64) Vector {
	return r.Origin.Add(r.Direction.Scale(t))
}
