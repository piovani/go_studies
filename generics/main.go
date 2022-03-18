package main

import "fmt"

func sum[N int64 | float64](m []N) N {
	var resultado N
	for _, numeroDoArray := range m {
		resultado += numeroDoArray
	}

	return resultado
}

func main() {
	si := []int64{1, 2, 5, 9, 12}
	sf := []float64{1.2, 2.5, 5.56, 9.67, 12.99}

	fmt.Println(si)
	fmt.Println(sf)

	fmt.Println(sum[int64](si))
	fmt.Println(sum[float64](sf))
}
