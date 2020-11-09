package geometry

// TODO should be in its own package but couldn't figure out the circular dependacies

// Material is an interface that defines what material an Element in the world is
type Material interface {
	Scatter(input Ray, hit HitRecord) (bool, Ray)
	Albedo() Vector
}
