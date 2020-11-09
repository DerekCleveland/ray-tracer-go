package geometry

// World defines a struct that consists of elements of the world
type World struct {
	Elements []Hitable
}

// CheckForHit implementation
func (w *World) CheckForHit(r *Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closest := tMax
	record := HitRecord{}

	for _, element := range w.Elements {
		hit, tempRecord := element.CheckForHit(r,tMin, closest)

		if hit {
			hitAnything = true
			closest = tempRecord.T
			record = tempRecord
		}
	}

	return hitAnything, record
}