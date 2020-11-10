package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"ray-tracer-go/pkg/camera"
	"ray-tracer-go/pkg/geometry"
	"ray-tracer-go/pkg/geometry/shape"
	"time"
)

// Image constants - NOT SURE IF I WANT TO USE THIS YET
// const (
// 	nx = 400
// 	ny = 200
// 	ns = 100
// 	c = 255.99
// )

func main() {
	log.Println("⏳ Starting ray tracer")

	err := rayTracer()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("⌛️ Finished ray tracer")
}

func rayTracer() error {
	// Define image filename
	filename := time.Now().Format("2006_01_02_150405") + ".PPM"
	path := "images/"
	filenameFullPath := path + filename

	// Check that directory exists. If not create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}

	// Create file
	f, err := os.Create(filenameFullPath)
	if err != nil {
		return err
	}
	defer f.Close()
	log.Println("Created file at", filenameFullPath)

	// Defines the resolution of the image
	nx := 400
	ny := 200
	// Number of samples we take for AA. The larger the value the smoother the transition but longer the processing time
	ns := 100
	// TODO determine what this value is for
	c := 255.99
	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))

	// Setup camera properties
	lookFrom := geometry.Vector{X: 3, Y: 3, Z: 2}
	lookAt := geometry.Vector{X: 0, Y: 0, Z: -1}
	vUp := geometry.Vector{X: 0, Y: 1, Z: 0}
	focusDist := lookFrom.Subtract(lookAt).Length()
	aperture := 2.0

	// Create a new camera
	camera := camera.NewCamera(lookFrom, lookAt, vUp, 20, float64(nx)/float64(ny), aperture, focusDist)

	// // Setup red and blue sphere
	// radius := math.Cos(math.Pi/4)

	// blueSphere := shape.Sphere{
	// 	Center: geometry.Vector{X: -radius, Y: 0, Z: -1},
	// 	Radius: radius,
	// 	Material: geometry.Lambertian{A: geometry.Vector{X: 0, Y: 0, Z: 1}},
	// }

	// redSphere := shape.Sphere{
	// 	Center: geometry.Vector{X: radius, Y: 0, Z: -1},
	// 	Radius: radius,
	// 	Material: geometry.Lambertian{A: geometry.Vector{X: 1, Y: 0, Z: 0}},
	// }
	// world := geometry.World{}

	// world.Add(&blueSphere)
	// world.Add(&redSphere)

	// Setup world
	// Brownish sphere
	// sphere := shape.Sphere{
	// 	Center:   geometry.Vector{X: 0, Y: 0, Z: -1},
	// 	Radius:   0.5,
	// 	Material: geometry.Lambertian{A: geometry.Vector{X: 0.8, Y: 0.3, Z: 0.3}},
	// }
	// Blue sphere
	sphere := shape.Sphere{
		Center:   geometry.Vector{X: 0, Y: 0, Z: -1},
		Radius:   0.5,
		Material: geometry.Lambertian{A: geometry.Vector{X: 0.1, Y: 0.2, Z: 0.5}},
	}

	floor := shape.Sphere{
		Center:   geometry.Vector{X: 0, Y: -100.5, Z: -1},
		Radius:   100,
		Material: geometry.Lambertian{A: geometry.Vector{X: 0.8, Y: 0.8, Z: 0.0}},
	}

	metalSphere1 := shape.Sphere{
		Center:   geometry.Vector{X: 1, Y: 0, Z: -1},
		Radius:   0.5,
		Material: geometry.Metal{A: geometry.Vector{X: 0.8, Y: 0.6, Z: 0.2}, Fuzz: 0.3},
	}

	// metalSphere2 := shape.Sphere{
	// 	Center:   geometry.Vector{X: -1, Y: 0, Z: -1},
	// 	Radius:   0.5,
	// 	Material: geometry.Metal{A: geometry.Vector{X: 0.8, Y: 0.8, Z: 0.8}, Fuzz: 0.3},
	// }

	dielectricSphere := shape.Sphere{
		Center:   geometry.Vector{X: -1, Y: 0, Z: -1},
		Radius:   0.5,
		Material: geometry.Dielectric{RefIndex: 1.5},
	}

	dielectricSphere2 := shape.Sphere{
		Center:   geometry.Vector{X: -1, Y: 0, Z: -1},
		Radius:   -0.45,
		Material: geometry.Dielectric{RefIndex: 1.5},
	}

	world := geometry.World{}

	world.Add(&sphere)
	world.Add(&floor)
	world.Add(&metalSphere1)
	//world.Add(&metalSphere2)
	world.Add(&dielectricSphere)
	world.Add(&dielectricSphere2)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			rgb := geometry.Vector{}

			// Sample rays for anti-aliasing
			for s := 0; s < ns; s++ {
				var u float64 = (float64(i) + rand.Float64()) / float64(nx)
				var v float64 = (float64(j) + rand.Float64()) / float64(ny)

				r := camera.RayAt(u, v)
				col := color(r, &world, 0)
				rgb = rgb.Add(col)
			}

			// Scale rgb vector by 1/ns
			rgb = rgb.Scale((1 / float64(ns)))
			// Take the square root of each column of rgb vector
			rgb = geometry.Vector{X: math.Sqrt(rgb.X), Y: math.Sqrt(rgb.Y), Z: math.Sqrt(rgb.Z)}

			// Color intensity
			ir := int(c * rgb.X)
			ig := int(c * rgb.Y)
			ib := int(c * rgb.Z)

			f.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	return nil
}

// color returns a vector...assumbably to change the color...but not sure how yet
func color(r geometry.Ray, world geometry.Hitable, depth int) geometry.Vector {
	hit, record := world.CheckForHit(r, 0.001, math.MaxFloat64)

	// if hit {
	// 	target := record.Point.Add(record.Normal).Add(RandomInUnitSphere())
	// 	return color(&geometry.Ray{Origin: record.Point, Direction: target.Subtract(record.P)}, world).Scale(0.5)
	// }

	if hit {
		if depth < 50 {
			scattered, scatteredRay := record.Scatter(r, record)
			if scattered {
				newColor := color(scatteredRay, world, depth+1)
				return record.Material.Albedo().Multiply(newColor)
			}
		}
	}

	// Make unit vector so y is between -1.0 and 1.0
	unitDirection := r.Direction.Normalize()

	var t float64 = 0.5 * (unitDirection.Y + 1.0)

	// The two vectors here are what creates the sky(Blue to white gradient of the background)
	return geometry.Vector{X: 1.0, Y: 1.0, Z: 1.0}.Scale(1.0 - t).Add(geometry.Vector{X: 0.5, Y: 0.7, Z: 1.0}.Scale(t))
}
