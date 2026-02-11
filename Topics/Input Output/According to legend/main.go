package main

import "fmt"

func main() {
	var system, moons, grid string
	var latitude, longitude float64
	var rotation, orbital int

	fmt.Scan(&system, &moons, &grid)
	fmt.Scan(&latitude, &longitude)
	fmt.Scan(&rotation, &orbital)

	fmt.Println("System:", system)
	fmt.Println("Moons:", moons)
	fmt.Println("Grid square:", grid)

	fmt.Println("\nLatitude:", latitude)
	fmt.Println("Longitude:", longitude)

	fmt.Println("\nRotation period in days:", rotation)
	fmt.Println("Orbital period in days:", orbital)
}
