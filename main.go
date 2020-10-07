package main

import (
	"fmt"
	"log"
	"os"
	"ray-tracer-go/pkg/geometry"
	"time"
)

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
	nx := 1280
	ny := 720
	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			vec := geometry.Vector{
				X: float64(i) / float64(nx),
				Y: float64(j) / float64(ny),
				Z: 0.2,
			}

			// Before vector implementation
			// r := float64(i) / float64(nx)
			// g := float64(j) / float64(ny)
			// b := 0.2

			ir := int(255.99 * vec.X)
			ig := int(255.99 * vec.Y)
			ib := int(255.99 * vec.Z)
			f.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	return nil
}
