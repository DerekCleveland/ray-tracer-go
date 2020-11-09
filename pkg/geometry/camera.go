package geometry

// Camera defines all the vectors needed for a camera position
type Camera struct {
	origin          Vector
	lowerLeftCorner Vector
	horizontal      Vector
	vertical        Vector
}

// TODO add ability to send in values to position camera at desired location

// NewCamera returns a camera position at the coded coordinates
func NewCamera() Camera {
	c := Camera{}

	c.origin = Vector{0.0, 0.0, 0.0}
	c.lowerLeftCorner = Vector{-2.0, -1.0, -1.0}
	c.horizontal = Vector{4.0, 0.0, 0.0}
	c.vertical = Vector{0.0, 2.0, 0.0}

	return c
}

func (c Camera) position(u float64, v float64) Vector {
	horizontal := c.horizontal.Scale(u)
	vertical := c.vertical.Scale(v)

	return horizontal.Add(vertical)
}

func (c Camera) direction(position Vector) Vector {
	return c.lowerLeftCorner.Add(position)
}

// RayAt returns the position of a ray for a given position
func (c Camera) RayAt(u float64, v float64) Ray {
	position := c.position(u, v)
	direction := c.direction(position)

	return Ray{
		Origin:    c.origin,
		Direction: direction,
	}
}
