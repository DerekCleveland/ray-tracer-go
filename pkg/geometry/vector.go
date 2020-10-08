package geometry

import (
	"fmt"
	"math"
)

// Vector is a struct defining an object that has both a magnitude and a direction
type Vector struct {
	X float64
	Y float64
	Z float64
}

// ToString takes a vector value receiver and returns it as a string
// TODO probably better to just use another image format that handles float64
func (v1 Vector) ToString() string {
	return fmt.Sprintf("%f %f %f", v1.X, v1.Y, v1.Z)
}

// Add takes in a vector as well as a vector value receiver and adds them together (v1 + v2)
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{v1.X + v2.X, v1.Y + v2.Y, v1.Z + v2.Z}
}

// Subtract takes in a vector as well as a vector value receiver and subtracts them from one another (v1 - v2)
func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{v1.X - v2.X, v1.Y - v2.Y, v1.Z - v2.Z}
}

// Multiply takes in a vector as well as a vector value receiver and multiplies the two (v1 * v2)
func (v1 Vector) Multiply(v2 Vector) Vector {
	return Vector{v1.X * v2.X, v1.Y * v2.Y, v1.Z * v2.Z}
}

// Divide takes in a vector as well as a vector value receiver and divides the two (v1 / v2)
func (v1 Vector) Divide(v2 Vector) Vector {
	return Vector{v1.X / v2.X, v1.Y / v2.Y, v1.Z / v2.Z}
}

// Dot takes in a vector as well as a vector value receiver and gets the dot product of the two
func (v1 Vector) Dot(v2 Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

// Cross takes in a vector as well as a vector value receiver and gets the cross product of the two
// TODO could of screwed this up a tad
func (v1 Vector) Cross(v2 Vector) Vector {
	return Vector{v1.Y*v2.Z - v1.Z*v2.Y, -(v1.X*v2.Z - v1.Z*v2.X), v1.X*v2.Y - v1.Y*v2.X}
}

// Scale takes in a value to scale the vector value receiver (v1 * s)
func (v1 Vector) Scale(s float64) Vector {
	return Vector{v1.X * s, v1.Y * s, v1.Z * s}
}

// Length returns back the length(also known as magnitude) of the vector value receiver
// TODO profling the two different ways of returns would be best
func (v1 Vector) Length() float64 {
	// return math.Sqrt(math.Pow(v1.X, 2) + math.Pow(v1.Y, 2) + math.Pow(v1.Z, 2))
	return math.Sqrt(v1.X*v1.X + v1.Y*v1.Y + v1.Z*v1.Z)
}

// Normalize returns back a vector that has been normalized(also known as converted to unit vectors) of the vector value receiver
func (v1 Vector) Normalize() Vector {
	inverse := 1.0 / v1.Length()
	return Vector{v1.X * inverse, v1.Y * inverse, v1.Z * inverse}
}
