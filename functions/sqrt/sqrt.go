package sqrt

import (
	"fmt"
	"math"
)

func Sqrt(x, initial float64, limit int) (float64, int) {
	result := initial
	i := 1

	for ; i <= limit; i++ {
		previous := result
		result -= (result*result - x) / (2 * result)
		difference := math.Abs(result - previous)

		if difference < 0.001 {
			break
		}
	}

	fmt.Printf("Square root of %f = %f, found in %d iterations with initial value %f\n", x, result, i, initial)

	return result, i
}
