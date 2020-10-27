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
	for i := 0; i < 2000000; i++ {
		// generate the random coordinates
		x = rand.Float64()
		y = rand.Float64()

		// if the coordinate is within the quarter circle
    	if (x*x + y*y <= 1.0) {
    		inside++
    	}
    	outside++
	}
	// this ratio will be pi, because of the area of the quarter circle
	fmt.Print("The algorithm predicts Pi to be: ")
	fmt.Println(4*(float64(inside)/float64(outside)))
}