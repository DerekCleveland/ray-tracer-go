package geometry

// HitRecord keeps track of a multitude of things for when a ray hits a point
type HitRecord struct {
	T      float64
	Point  Vector
	Normal Vector
	Material
}

// Hitable requires that geometry be Hit and return hit properties
type Hitable interface {
	CheckForHit(r Ray, tMin float64, tmax float64) (bool, HitRecord)
}
