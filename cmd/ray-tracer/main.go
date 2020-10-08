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
	nx := 200
	ny := 100
	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))

	lowerLeftCorner := geometry.Vector{X: -2.0, Y: -1.0, Z: -1.0}
	horizontal := geometry.Vector{X: 4.0, Y: 0.0, Z: 0.0}
	vertical := geometry.Vector{X: 0.0, Y: 2.0, Z: 0.0}
	origin := geometry.Vector{X: 0.0, Y: 0.0, Z: 0.0}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			var u float64 = float64(i) / float64(nx)
			var v float64 = float64(j) / float64(ny)

			product := lowerLeftCorner.Add(horizontal.Scale(u)).Add(vertical.Scale(v))

			r := geometry.Ray{
				Origin: origin,
				Direction: product,
			}

			col := r.Color()

			ir := int(255.99 * col.X)
			ig := int(255.99 * col.Y)
			ib := int(255.99 * col.Z)

			f.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	return nil
}
