package network

import (
	"fmt"
	"math"
	"nnet/matrix"
)

type neuralNetwork struct {
	iNodes int
	hNodes int
	oNodes int
	wih    *matrix.Matrix //weight input-hidden layer
	who    *matrix.Matrix //weight hidden-output layer
	lr     float64
}

type NeuralNetwork interface {
	Train(input, target []float64)
	Query(input []float64) []float64
}

func New(inputNodes, hiddenNodes, outputNodes int, learnIngrate float64) NeuralNetwork {
	wih := matrix.NewRand(hiddenNodes, inputNodes)
	who := matrix.NewRand(outputNodes, hiddenNodes)
	return &neuralNetwork{
		iNodes: inputNodes,
		hNodes: hiddenNodes,
		oNodes: outputNodes,
		wih:    wih,
		who:    who,
		lr:     learnIngrate,
	}
}

func (n *neuralNetwork) Train(input, target []float64) {

	in := matrix.NewEmpty(len(input), 1)
	in.Vals = input
	targ := matrix.NewEmpty(len(target), 1)
	targ.Vals = target
	//входные сигналы для скрытого слоя
	hInput := n.wih.Mult(in)
	//выходные сигналы скрытого слоя
	hOutput := hInput.Map(sigmoid)

	//входные сигналы выходного слоя
	fInput := n.who.Mult(hOutput)
	//выходные сигналы выходного слоя
	fOutput := fInput.Map(sigmoid)

	//выходная ошибка
	outputErrors := targ.Sub(fOutput)
	//hiddenErrors
	hiddenErrors := n.who.Transpose().Mult(outputErrors)
	//Поправочные коэффициенты
	fmt.Println("Input-hidden before correct:")
	fmt.Println(n.wih)
	fmt.Println("Hidden-output before correct:")
	fmt.Println(n.who)
	n.who.Vals = n.who.Sum(
		fOutput.Map(func(x float64) float64 { return 1.0 - x }).
			MultByEl(fOutput).
			MultByEl(outputErrors).
			Mult(hOutput.Transpose()).
			Map(func(x float64) float64 { return x * n.lr })).Vals
	n.wih.Vals = n.wih.Sum(
		hOutput.Map(func(x float64) float64 { return 1.0 - x }).
			MultByEl(hOutput).
			MultByEl(hiddenErrors).
			Mult(in.Transpose()).
			Map(func(x float64) float64 { return x * n.lr })).Vals
	fmt.Println("Input-hidden after correct:")
	fmt.Println(n.wih)
	fmt.Println("Hidden-output after correct:")
	fmt.Println(n.who)
}

func (n *neuralNetwork) Query(input []float64) []float64 {

	in := matrix.NewEmpty(len(input), 1)
	in.Vals = input
	//входные сигналы для скрытого слоя
	hInput := n.wih.Mult(in)
	//выходные сигналы скрытого слоя
	hOutput := hInput.Map(sigmoid)

	//входные сигналы выходного слоя
	fInput := n.who.Mult(hOutput)
	//выходные сигналы выходного слоя
	fOutput := fInput.Map(sigmoid)

	return fOutput.Vals
}

// sigmoid implements the sigmoid function
// for use in activation functions.
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
