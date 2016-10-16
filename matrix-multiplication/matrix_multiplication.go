package matrixmultiplication

import "errors"

type Matrix struct {
	data       [][]int
	rows, cols int
}

func New(data [][]int) (*Matrix, error) {
	if data == nil {
		return nil, errors.New("Not a Matrix")
	}
	if len(data) == 0 || len(data[0]) == 0 {
		return nil, errors.New("one Dimension is Zero, not a Matrix")
	}
	rowLength := len(data[0])
	for _, r := range data {
		if rowLength != len(r) {
			return nil, errors.New("Dimensions rows aren't event")
		}
	}
	return &Matrix{data, len(data), len(data[0])}, nil
}

func (m *Matrix) Multiply(m2 *Matrix) ([][]int, error) {
	if m2 == nil {
		return nil, errors.New("Matrices can't be nil")
	}
	matrix1 := m.data
	matrix2 := m2.data
	rows := len(matrix1)
	cols := len(matrix2[0])
	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("Columns of first Matrix need to be the same as Rows of second for multiplication")
	}
	result := make([][]int, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[i][j] = 0
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result, nil
}
