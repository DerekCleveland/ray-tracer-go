package geometry


// HitRecord keeps track of TODO finish this
type HitRecord struct {
	T float64
	P Vector
	Normal Vector
}

// Hitable requires that geometry be Hit and return hit properties
type Hitable interface {
	CheckForHit(r *Ray, tMin float64, tmax float64) (bool, HitRecord)
}