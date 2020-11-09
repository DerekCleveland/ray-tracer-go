package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
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
	nx := 200
	ny := 100
	// Number of samples we take for AA. The larger the value the smoother the transition but longer the processing time
	ns := 100
	// TODO determine what this value is for
	c := 255.99
	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))

	// Create a new camera
	camera := geometry.NewCamera()

	// Setup world
	sphere := shape.Sphere{
		Center: geometry.Vector{X: 0, Y: 0, Z: -1},
		Radius: 0.5,
	}

	floor := shape.Sphere{
		Center: geometry.Vector{X: 0, Y: -100.5, Z: -1},
		Radius: 100,
	}

	world := geometry.World{
		Elements: []geometry.Hitable{&sphere, &floor},
	}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			rgb := geometry.Vector{}

			// Sample rays for anti-aliasing
			for s := 0; s < ns; s++ {
				var u float64 = (float64(i) + rand.Float64()) / float64(nx)
				var v float64 = (float64(j) + rand.Float64()) / float64(ny)

				r := camera.RayAt(u, v)
				// col := r.Color()
				col := color(&r, &world)
				rgb = rgb.Add(col)
			}

			// TODO this might be wrong? I think my math checks out though
			rgb = rgb.Scale((1/float64(ns)))

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
func color(r *geometry.Ray, h geometry.Hitable) geometry.Vector {
	hit, record := h.CheckForHit(r, 0.0, math.MaxFloat64)

	if hit {
		return geometry.Vector{X: record.Normal.X + 1, Y: record.Normal.Y + 1, Z: record.Normal.Z + 1}.Scale(0.5) 
	}

	// Make unit vector so y is between -1.0 and 1.0
	unitDirection := r.Direction.Normalize()

	var t float64 = 0.5 * (unitDirection.Y + 1.0)

	// The two vectors here are what creates the sky(Blue to white gradient of the background)
	return geometry.Vector{X: 1.0, Y: 1.0, Z: 1.0}.Scale(1.0 - t).Add(geometry.Vector{X: 0.5, Y: 0.7, Z: 1.0}.Scale(t))
}
