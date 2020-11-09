package geometry

// World defines a struct that consists of elements of the world
type World struct {
	elements []Hitable
}

// Add allows you to add elements to the world
func (w *World) Add(h Hitable) {
	w.elements = append(w.elements, h)
}

// CheckForHit implementation
func (w *World) CheckForHit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closest := tMax
	record := HitRecord{}

	for _, element := range w.elements {
		hit, tempRecord := element.CheckForHit(r, tMin, closest)

		if hit {
			hitAnything = true
			closest = tempRecord.T
			record = tempRecord
		}
	}

	return hitAnything, record
}
