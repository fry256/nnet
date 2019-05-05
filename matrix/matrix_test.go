package matrix

import (
	"reflect"
	"testing"
)

func TestNewEmpty(t *testing.T) {
	type args struct {
		rows int
		cols int
	}
	tests := []struct {
		name string
		args args
		want *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmpty(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Mult(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		Vals []float64
	}
	type args struct {
		b *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				Vals: tt.fields.Vals,
			}
			if got := a.Mult(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_MultByEl(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		Vals []float64
	}
	type args struct {
		b *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				Vals: tt.fields.Vals,
			}
			if got := a.MultByEl(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.MultByEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Transpose(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		Vals []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				Vals: tt.fields.Vals,
			}
			if got := a.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Map(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		Vals []float64
	}
	type args struct {
		f func(x float64) float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				Vals: tt.fields.Vals,
			}
			if got := a.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Sum(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		Vals []float64
	}
	type args struct {
		b *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				Vals: tt.fields.Vals,
			}
			if got := a.Sum(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Sub(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		Vals []float64
	}
	type args struct {
		b *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Matrix
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				Vals: tt.fields.Vals,
			}
			if got := a.Sub(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Matrix.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_String(t *testing.T) {
	type fields struct {
		Rows int
		Cols int
		Vals []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Matrix{
				Rows: tt.fields.Rows,
				Cols: tt.fields.Cols,
				Vals: tt.fields.Vals,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("Matrix.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
