package camera

import (
	"math"
	"math/rand"
	"ray-tracer-go/pkg/geometry"
)

// Camera defines all the vectors needed for a camera position
type Camera struct {
	origin          geometry.Vector
	lowerLeftCorner geometry.Vector
	horizontal      geometry.Vector
	vertical        geometry.Vector
	u               geometry.Vector
	v               geometry.Vector
	w               geometry.Vector
	lensRadius      float64
}

// NewCamera returns a camera position at the coded coordinates
func NewCamera(lookFrom geometry.Vector, lookAt geometry.Vector, vUp geometry.Vector, vfov float64, aspect float64, aperture float64, focusDist float64) Camera {
	c := Camera{}

	c.origin = lookFrom
	c.lensRadius = aperture / 2

	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := lookFrom.Subtract(lookAt).Normalize()
	u := vUp.Cross(w).Normalize()
	v := w.Cross(u)

	x := u.Scale(halfWidth * focusDist)
	y := v.Scale(halfHeight * focusDist)

	c.lowerLeftCorner = c.origin.Subtract(x).Subtract(y).Subtract(w.Scale(focusDist))
	c.horizontal = x.Scale(2)
	c.vertical = y.Scale(2)

	c.w = w
	c.u = u
	c.v = v

	return c
}

// RayAt returns the position of a ray for a given position
func (c Camera) RayAt(s float64, t float64) geometry.Ray {
	rd := randomInUnitDisc().Scale(c.lensRadius)
	offset := c.u.Scale(rd.X).Add(c.v.Scale(rd.Y))

	horizontal := c.horizontal.Scale(s)
	vertical := c.vertical.Scale(t)

	origin := c.origin.Add(offset)
	direction := c.lowerLeftCorner.Add(horizontal).Add(vertical).Subtract(origin)

	return geometry.Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func randomInUnitDisc() geometry.Vector {
	var p geometry.Vector
	for {
		p = geometry.Vector{X: rand.Float64(), Y: rand.Float64(), Z: 0}.Scale(2).Subtract(geometry.Vector{X: 1, Y: 1, Z: 0})
		if p.Dot(p) < 1.0 {
			return p
		}
	}
}
