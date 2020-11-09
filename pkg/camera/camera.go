package camera

import "ray-tracer-go/pkg/geometry"

// Camera defines all the vectors needed for a camera position
type Camera struct {
	origin          geometry.Vector
	lowerLeftCorner geometry.Vector
	horizontal      geometry.Vector
	vertical        geometry.Vector
}

// TODO add ability to send in values to position camera at desired location

// NewCamera returns a camera position at the coded coordinates
func NewCamera() Camera {
	c := Camera{}

	c.origin = geometry.Vector{0.0, 0.0, 0.0}
	c.lowerLeftCorner = geometry.Vector{-2.0, -1.0, -1.0}
	c.horizontal = geometry.Vector{4.0, 0.0, 0.0}
	c.vertical = geometry.Vector{0.0, 2.0, 0.0}

	return c
}

func (c Camera) position(u float64, v float64) geometry.Vector {
	horizontal := c.horizontal.Scale(u)
	vertical := c.vertical.Scale(v)

	return horizontal.Add(vertical)
}

func (c Camera) direction(position geometry.Vector) geometry.Vector {
	return c.lowerLeftCorner.Add(position)
}

// RayAt returns the position of a ray for a given position
func (c Camera) RayAt(u float64, v float64) geometry.Ray {
	position := c.position(u, v)
	direction := c.direction(position)

	return geometry.Ray{
		Origin:    c.origin,
		Direction: direction,
	}
}
