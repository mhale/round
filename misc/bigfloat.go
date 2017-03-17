package main

import "fmt"
import "math/big"

func main() {
	values := []float64{5.5, 2.5, 1.6, 1.1, 1.0, -1.0, -1.1, -1.6, -2.5, -5.5}

	fmt.Print("   x")
	for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
			fmt.Printf("  %s", mode)
	}
	fmt.Println()

	for _, x := range values {
		fmt.Printf("%4.1f", x)
		for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
			f := new(big.Float).SetMode(mode).SetFloat64(x)
			if x > 4.0 || x < -4.0 {
				f.SetPrec(3)
			} else if x > 2.0 || x < -2.0 {
				f.SetPrec(2)
			} else {
				f.SetPrec(1)
			}
			fmt.Printf("  %*.1f", len(mode.String()), f)
		}
		fmt.Println()
	}
}
