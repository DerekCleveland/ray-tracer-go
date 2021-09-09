package shape

import (
	"math"
	"ray-tracer-go/pkg/geometry"
)

// Sphere defines a sphere object in 3d space that implements shape and has a center and radius
type Sphere struct {
	Center geometry.Vector
	Radius float64
	geometry.Material
}

// CheckForHit checks if the passed in ray hits the sphere
func (sphere *Sphere) CheckForHit(ray geometry.Ray, tMin float64, tMax float64) (bool, geometry.HitRecord) {
	// Improve variables names and comments
	oc := ray.Origin.Subtract(sphere.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := oc.Dot(ray.Direction)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius

	discriminant := b*b - a*c

	rec := geometry.HitRecord{Material: sphere.Material}

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			// TODO might need to implement pointAtParameter...not sure if I just improved the name or what
			rec.Point = ray.PointOnRay(rec.T)
			rec.Normal = (rec.Point.Subtract(sphere.Center)).Scale(1 / sphere.Radius)
			return true, rec
		}

		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			rec.T = temp
			rec.Point = ray.PointOnRay(rec.T)
			rec.Normal = (rec.Point.Subtract(sphere.Center)).Scale(1 / sphere.Radius)
			return true, rec
		}
	}

	return false, geometry.HitRecord{}
}
