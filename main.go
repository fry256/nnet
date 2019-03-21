package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type neuralNetwork struct {
	iNodes int
	hNodes int
	oNodes int
	wih    [][]float64 //weight input-hidden layer
	who    [][]float64 //weight hidden-output layer
	lr     float64
}

type NeuralNetwork interface {
	train(input, target [][]float64)
	query(input [][]float64) [][]float64
}

func New(inputNodes, hiddenNodes, outputNodes int, learnIngrate float64) NeuralNetwork {
	wih := randMatrix(hiddenNodes, inputNodes)
	who := randMatrix(outputNodes, hiddenNodes)
	return &neuralNetwork{
		iNodes: inputNodes,
		hNodes: hiddenNodes,
		oNodes: outputNodes,
		wih:    wih,
		who:    who,
		lr:     learnIngrate,
	}
}
func (n *neuralNetwork) train(input, target [][]float64) {

	//входные сигналы для скрытого слоя
	hInput, err := dot(n.wih, input)
	if err != nil {
		fmt.Println(err)
	}
	//выходные сигналы скрытого слоя
	hOutput := activation(hInput)

	//входные сигналы выходного слоя
	fInput, err := dot(n.who, hOutput)
	if err != nil {
		fmt.Println(err)
	}
	//выходные сигналы выходного слоя
	fOutput := activation(fInput)

	//выходная ошибка
	outputErrors := make([][]float64, len(target))
	for i := range target {
		for j := range target[i] {
			outputErrors[i] = append(outputErrors[i], target[i][j]-fOutput[i][j])
		}
	}
	//hiddenErrors
	_, err = dot(transp(n.who), outputErrors)
	if err != nil {
		fmt.Println(err)
	}
	//Поправочные коэффициенты
	a := sub(1, fOutput)
	b := increaseMatrix(fOutput, a)
	c := increaseMatrix(outputErrors, b)
	d, err := dot(c, transp(hOutput))
	if err != nil {
		fmt.Println(err)
	}
	e := increaseEvery(n.lr, d)
	fmt.Println(sumMatrix(n.who, e))
	fmt.Println(n.who)
	return
}
func (n *neuralNetwork) query(input [][]float64) [][]float64 {

	//входные сигналы для скрытого слоя
	hInput, err := dot(n.wih, input)
	if err != nil {
		fmt.Println(err)
	}
	//выходные сигналы скрытого слоя
	hOutput := activation(hInput)

	//входные сигналы выходного слоя
	fInput, err := dot(n.who, hOutput)
	if err != nil {
		fmt.Println(err)
	}
	//выходные сигналы выходного слоя
	fOutput := activation(fInput)

	return fOutput
}

func main() {
	i := 3
	h := 3
	o := 3
	lr := 0.3
	n := New(i, h, o, lr)
	/*result := n.query([][]float64{
		[]float64{1.0},
		[]float64{0.5},
		[]float64{-1.5},
	})

	fmt.Println(result)*/

	input := [][]float64{
		[]float64{1.0},
		[]float64{0.5},
		[]float64{-1.5},
	}

	target := [][]float64{
		[]float64{0.01},
		[]float64{0.99},
		[]float64{0.01},
	}
	n.train(input, target)

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

//Создание случайной матрицы
func randMatrix(rows, columns int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			matrix[i] = append(matrix[i], rand.Float64()-0.5)
		}
	}
	return matrix
}

//                | 2 1 |
//| 2 3 5 2 5 |   | 1 3 |   | 40 37 |
//| 2 4 7 1 3 | x | 4 2 | = | 43 37 |
//| 1 5 4 8 3 |   | 4 3 |   | 58 54 |
//                | 1 2 |
//
//c[1][1] = a[1][1] * b[1][1] + a[1][2] * b[2][1] + a[1][3] * b[3][1] + a[1][4] * b[4][1] + a[1][5] * b[5][1]
//
//
//
// c[1][1] = 2 * 2 + 3 * 1 + 5 * 4 + 2 * 4 + 5 * 1 = 4 + 3 + 20 + 8 + 5 = 40
// c[2][1] = 2 * 2 + 4 * 1 + 7 * 4 + 1 * 4 + 3 * 1 = 4 + 4 + 28 + 4 + 3 = 43
// c[3][1] = 1 * 2 + 5 * 1 + 4 * 4 + 8 * 4 + 3 * 1 = 2 + 5 + 16 + 32 + 3 = 58
// c[1][2] = 2 * 1 + 3 * 3 + 5 * 2 + 2 * 3 + 5 * 2 = 2 + 9 + 10 + 6 + 10 = 37
// c[2][2] = 2 * 1 + 4 * 3 + 7 * 2 + 1 * 3 + 3 * 2 = 2 + 12 + 14 + 3 + 6 = 37
// c[3][2] = 1 * 1 + 5 * 3 + 4 * 2 + 8 * 3 + 3 * 2 = 1 + 15 + 8 + 24 + 6 = 54
//Умножение матриц
func dot(a [][]float64, b [][]float64) ([][]float64, error) {
	if len(a) == 0 || len(b) == 0 {
		return nil, errors.New("Error! Empty...")
	}
	if len(a[0]) != len(b) {
		return nil, errors.New("Error!")
	}

	var item float64
	c := make([][]float64, len(a))
	for i := 0; i < len(b[0]); i++ {
		for j := 0; j < len(a); j++ {
			item = 0.0
			for k := 0; k < len(a[j]); k++ {
				item += a[j][k] * b[k][i]
			}
			c[j] = append(c[j], item)
		}
	}

	return c, nil
}

// sigmoid implements the sigmoid function
// for use in activation functions.
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func activation(input [][]float64) [][]float64 {
	result := make([][]float64, len(input))

	for i := range input {
		for j := range input[i] {
			result[i] = append(result[i], sigmoid(input[i][j]))
		}
	}

	return result
}

// | 1.1, 2.1, 4.5 |    | 1.1, 4.3, 0.7 |T
// | 4.3, 0.6, 0.3 | == | 2.1, 0.6, 0.9 |
// | 0.7, 0.9, 1.7 |    | 4.5, 0.3, 1.7 |
//
// | 1 |             T
// | 4 | == | 1 4 2 |
// | 2 |
func transp(matrix [][]float64) [][]float64 {
	t := make([][]float64, len(matrix[0]))

	for i := range t {
		t[i] = make([]float64, len(matrix))
	}

	for i, row := range matrix {
		for j, col := range row {
			t[j][i] = col
		}
	}
	return t
}

func increaseMatrix(a, b [][]float64) [][]float64 {
	c := make([][]float64, len(a))

	for i := range a {
		c[i] = append(c[i], a[i][0]*b[i][0])
	}

	return c
}

func increaseFloat(a float64, b [][]float64) [][]float64 {
	c := make([][]float64, len(b))

	for i := range b {
		c[i] = append(c[i], a*b[i][0])
	}

	return c
}

func increaseEvery(a float64, b [][]float64) [][]float64 {
	c := make([][]float64, len(b))

	for i := range b {
		for j := range b[i] {
			c[i] = append(c[i], a*b[i][j])
		}
	}

	return c
}

func sumMatrix(a, b [][]float64) [][]float64 {
	c := make([][]float64, len(b))

	for i := range b {
		for j := range b[i] {
			c[i] = append(c[i], a[i][j]+b[i][j])
		}
	}

	return c
}

func sub(a float64, b [][]float64) [][]float64 {
	c := make([][]float64, len(b))

	for i := range b {
		c[i] = append(c[i], a-b[i][0])
	}

	return c
}
