// gotests1 project main.go
package main

import "fmt"
import "mh"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
