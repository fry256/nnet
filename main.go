package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"nnet/matrix"
	"nnet/network"
	"os"
	"strconv"
)

func main() {
	i := 784
	h := 200
	o := 10
	lr := 0.2
	n := network.New(i, h, o, lr)
	epochs := 5
	for i := 0; i < epochs; i++ {
		csvTrain, err := os.Open("mnist_train_100.csv")
		if err != nil {
			log.Fatal(err)
		}
		r := csv.NewReader(bufio.NewReader(csvTrain))

		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			correctLabel, err := strconv.Atoi(record[0])
			if err != nil {
				log.Fatal(err)
			}

			m := matrix.NewEmpty(28, 28)
			for key, val := range record[1:] {
				m.Vals[key], err = strconv.ParseFloat(val, 64)
				if err != nil {
					log.Fatal(err)
				}
			}

			target := []float64{0.01, 0.01, 0.01, 0.01, 0.01, 0.01, 0.01, 0.01, 0.01, 0.01}
			target[correctLabel] = 0.99

			input := convertRange(m)
			n.Train(input.Vals, target)
		}
	}

	csvTest, err := os.Open("mnist_test_10.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(bufio.NewReader(csvTest))
	countOfCorrectResults := 0

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		correctLabel, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		m := matrix.NewEmpty(28, 28)
		for key, val := range record[1:] {
			m.Vals[key], err = strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}
		}

		input := convertRange(m)

		output := n.Query(input.Vals)
		var label int
		max := output[0]

		for i, val := range output {
			if val > max {
				max = val
				label = i
			}
		}

		if label == correctLabel {
			countOfCorrectResults++
		}

		fmt.Printf("Label: %d, Result: %d\n", correctLabel, label)
	}

	fmt.Printf("Количество правильных ответов: %d\n", countOfCorrectResults)

}

func convertRange(m *matrix.Matrix) *matrix.Matrix {
	return m.Map(func(x float64) float64 {
		return x / 255.0
	}).Map(func(x float64) float64 {
		return x * 0.99
	}).Map(func(x float64) float64 {
		return x + 0.01
	})
}
