package main

import (
    "fmt"
    "math/rand"
)

func main() {
	// initialize variables for number of coordinates within/outside the quarter circle
	var inside int = 0
	var outside int = 0
	var x float64
	var y float64
	for i := 0; i < 1000000; i++ {
		// generate the random coordinates
		x = rand.Float64()
		y = rand.Float64()

		// if the coordinate is within the quarter circle
		fmt.Print(x, y)
    	fmt.Println()

    	if (x*x + y*y <= 1) {
    		inside++
    	}
    	outside++
	}
}