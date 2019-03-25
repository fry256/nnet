package matrix

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Matrix struct {
	Rows, Cols int
	Vals       []float64
}

func NewEmpty(rows, cols int) *Matrix {
	a := new(Matrix)

	a.Rows = rows
	a.Cols = cols
	a.Vals = make([]float64, rows*cols)

	return a
}

//Создание случайной матрицы
func NewRand(rows, cols int) *Matrix {
	a := NewEmpty(rows, cols)

	for i := range a.Vals {
		a.Vals[i] = rand.Float64()
	}

	return a
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
func (a *Matrix) Mult(b *Matrix) *Matrix {
	if len(a.Vals) == 0 || len(b.Vals) == 0 {
		panic("Error! Empty...")
	}
	if a.Cols != b.Rows {
		panic("Error!")
	}

	c := NewEmpty(a.Rows, b.Cols)

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

func (a *Matrix) MultByEl(b *Matrix) *Matrix {
	c := NewEmpty(a.Cols, a.Rows)

	for i := range a.Vals {
		c.Vals[i] = a.Vals[i] * b.Vals[i]
	}

	return c
}

// | 1.1, 2.1, 4.5 |    | 1.1, 4.3, 0.7 |T
// | 4.3, 0.6, 0.3 | == | 2.1, 0.6, 0.9 |
// | 0.7, 0.9, 1.7 |    | 4.5, 0.3, 1.7 |
//
// | 1 |             T
// | 4 | == | 1 4 2 |
// | 2 |
func (a *Matrix) Transpose() *Matrix {
	b := NewEmpty(a.Cols, a.Rows)

	for i := 1; i <= a.Rows; i++ {
		for j := 1; j <= a.Cols; j++ {
			b.Vals[(j-1)*b.Cols+i-1] = a.Vals[(i-1)*a.Cols+j-1]
		}
	}

	return b
}

func (a *Matrix) Map(f func(x float64) float64) *Matrix {
	b := NewEmpty(a.Rows, a.Cols)

	for i, v := range a.Vals {
		b.Vals[i] = f(v)
	}

	return b
}

func (a *Matrix) Sum(b *Matrix) *Matrix {
	c := NewEmpty(a.Rows, a.Cols)

	for i := range a.Vals {
		c.Vals[i] = a.Vals[i] + b.Vals[i]
	}

	return c
}

func (a *Matrix) Sub(b *Matrix) *Matrix {
	c := NewEmpty(a.Rows, a.Cols)

	for i := range a.Vals {
		c.Vals[i] = a.Vals[i] - b.Vals[i]
	}

	return c
}

func (a *Matrix) String() string {
	buffer := new(bytes.Buffer)

	for i, elem := range a.Vals {
		buffer.WriteString(fmt.Sprintf(" %.3f ", elem))

		if (i+1)%a.Cols == 0 && i+1 != len(a.Vals) {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}
