package sqrt

import (
	"fmt"
)

func Demo() {
	values := []float64{1.0, 20.0, 67.3, 526544.0, 0.000324}

	for i := 0; i < len(values); i++ {
		fmt.Printf("Value %f: \n", values[i])

		Sqrt(values[i], 1.0, 50)
		Sqrt(values[i], values[i]/2, 50)

		fmt.Println()
	}
}
