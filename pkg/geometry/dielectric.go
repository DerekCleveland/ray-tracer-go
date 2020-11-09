package geometry

import (
	"math"
	"math/rand"
)

// Dielectric describes a dielectric material
type Dielectric struct {
	RefIndex float64
}

// Albedo implementation
func (d Dielectric) Albedo() Vector {
	// If changed to {1.0, 1.0, 0.0} will kill the blue channel. Can change between the two
	// to see the difference
	return Vector{X: 1.0, Y: 1.0, Z: 1.0}
}

// Scatter implementation
func (d Dielectric) Scatter(input Ray, hit HitRecord) (bool, Ray) {
	var outwardNormal Vector
	var niOverNt float64
	var cosine float64

	// TODO if something is off...its most likely in here that I screwed up
	if input.Direction.Dot(hit.Normal) > 0 {
		outwardNormal = hit.Normal.Scale(-1)
		niOverNt = d.RefIndex
		cosine = d.RefIndex * input.Direction.Dot(hit.Normal) / input.Direction.Length()
	} else {
		outwardNormal = hit.Normal
		niOverNt = 1.0 / d.RefIndex
		cosine = -input.Direction.Dot(hit.Normal) / input.Direction.Length()
	}

	var success bool
	var refracted Vector
	var reflectProbability float64

	if success, refracted = input.Direction.Refract(outwardNormal, niOverNt); success {
		reflectProbability = d.Schlick(cosine)
	} else {
		reflectProbability = 1.0
	}

	// Ray was reflected
	if rand.Float64() < reflectProbability {
		reflected := input.Direction.Reflect(hit.Normal)
		return true, Ray{hit.Point, reflected}
	}

	// Ray was refracted
	return true, Ray{hit.Point, refracted}
}

// Schlick is a polynomial approximation of glass reflectivity
func (d Dielectric) Schlick(cosine float64) float64 {
	var r0 float64 = (1 - d.RefIndex) / (1 + d.RefIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
