package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cube struct {
	x float64
}

// Implementing Shape3D interface
func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

// Implementing Shape3D interface
func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

type Sphere struct {
	r float64
}

// Implementing Shape3D interface
func (c Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * c.r * c.r * c.r
}

type shapes []Shape3D

// Implementing sort.Interface
func (a shapes) Len() int {
	return len(a)
}
func (a shapes) Less(i, j int) bool {
	return a[i].Vol() < a[j].Vol()
}
func (a shapes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func PrintShapes(a shapes) {
	for _, v := range a {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown data type!")
		}
	}
	fmt.Println()
}

func main() {
	data := shapes{}
	rand.New(rand.NewSource(42))

	minVal := 1.
	maxVal := 5.
	for i := 0; i < 3; i++ {
		cube := Cube{rF64(minVal, maxVal)}
		data = append(data, cube)
		
		cuboid := Cuboid{rF64(minVal, maxVal), rF64(minVal, maxVal), rF64(minVal, maxVal)}
		data = append(data, cuboid)
		
		sphere := Sphere{rF64(minVal, maxVal)}
		data = append(data, sphere)
	}

	// Cube: volume 19.91
	// Cuboid: volume 19.76
	// Sphere: volume 90.96
	// Cube: volume 17.31
	// Cuboid: volume 50.97
	// Sphere: volume 7.46
	// Cube: volume 78.51
	// Cuboid: volume 17.45
	// Sphere: volume 5.19
	PrintShapes(data)

	// Sphere: volume 5.19
	// Sphere: volume 7.46
	// Cube: volume 17.31
	// Cuboid: volume 17.45
	// Cuboid: volume 19.76
	// Cube: volume 19.91
	// Cuboid: volume 50.97
	// Cube: volume 78.51
	// Sphere: volume 90.96
	sort.Sort(data)
	PrintShapes(data)
}
