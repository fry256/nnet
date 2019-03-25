package main

import (
	"fmt"
	"nnet/matrix"
	"nnet/network"
)

func main() {
	i := 3
	h := 3
	o := 3
	lr := 0.3
	n := network.New(i, h, o, lr)
	/*result := n.Query([]float64{1.0, 0.5, -1.5})

	fmt.Println(result)*/

	input := []float64{1.0, 0.5, -1.5}
	target := []float64{0.01, 0.99, 0.01}
	n.Train(input, target)

	m1 := matrix.NewEmpty(3, 5)
	m1.Vals = []float64{2.0, 3.0, 5.0, 2.0, 5.0, 2.0, 4.0, 7.0, 1.0, 3.0, 1.0, 5.0, 4.0, 8.0, 3.0}

	m2 := matrix.NewEmpty(5, 2)
	m2.Vals = []float64{2.0, 1.0, 1.0, 3.0, 4.0, 2.0, 4.0, 3.0, 1.0, 2.0}

	fmt.Printf("\n\n\n*************************\n\n\n")
	fmt.Println(m1.Mult(m2))

	/*m1 := [][]float64{
		[]float64{2.0, 3.0, 5.0, 2.0, 5.0},
		[]float64{2.0, 4.0, 7.0, 1.0, 3.0},
		[]float64{1.0, 5.0, 4.0, 8.0, 3.0},
	}

	m2 := [][]float64{
		[]float64{2.0, 1.0},
		[]float64{1.0, 3.0},
		[]float64{4.0, 2.0},
		[]float64{4.0, 3.0},
		[]float64{1.0, 2.0},
	}

	m3 := make([][]float64, 0)

	m4 := [][]float64{
		[]float64{2.0},
		[]float64{1.0},
		[]float64{3.0},
	}

	r1, err := dot(m1, m2)
	if err != nil {
		fmt.Println(err)
	}
	r2, err := dot(m1, m3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(transp(m4))*/
}
