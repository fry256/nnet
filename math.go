package matrix

import (
	"errors"
	"math/rand"
)

type Matrix struct {
	Rows, Cols int
	Vals       []float64
}

func newEmpty(rows, cols int) *Matrix {
	a := new(Matrix)

	a.Rows = rows
	a.Cols = cols
	a.Vals = make([]float64, rows*cols)

	return a
}

//Создание случайной матрицы
func NewRand(rows, columns int) *Matrix {
	a := newEmpty(rows, cols)

	for i := range a.Vals {
		a.Vals[i] = rand.Float64()
	}

	return a
}

func (a *Matrix) Mult(b *Matrix) *Matrix {
	if len(a.Vals) == 0 || len(b.Vals) == 0 {
		return nil, errors.New("Error! Empty...")
	}
	if a.Cols != b.Rows {
		return nil, errors.New("Error!")
	}

	c := newEmpty(a.Rows, b.Cols)

	for i := 0; i < c.Rows; i++ {
		for j := 0; j < c.Cols; j++ {
			sum := float64(0)

			for k := 0; k < a.Cols; k++ {
				sum += a.Vals[i*a.Cols+k] * b.Vals[k*b.Cols+j]
			}

			c.Vals[i*c.Cols+j] = sum
		}
	}

	return c
}

func (a *Matrix) transpose() *Matrix {
	b := newEmpty(a.Cols, a.Rows)

	for i := 1; i <= a.Rows; i++ {
		for j := 1; j <= a.Cols; j++ {
			b.Vals[(j-1)*b.Cols+i-1] = a.Vals[(i-1)*a.Cols+j-1]
		}
	}

	return b
}

func (a *Matrix) Map(f func(x float64) float64) *Matrix {
	b := newEmpty(a.Rows, a.Cols)

	for i, v := range a.Vals {
		b[i] = f(v)
	}

	return b
}
