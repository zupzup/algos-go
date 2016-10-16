package matrixmultiplication

import (
	"fmt"
	"reflect"
	"testing"
)

var testsNew = []struct {
	description string
	in          [][]int
	ok          bool
	rows, cols  int
}{
	{
		"basic2x2",
		[][]int{
			{1, 2},
			{2, 3},
		},
		true,
		2,
		2,
	},
	{
		"basic1x2",
		[][]int{
			{1, 2},
		},
		true,
		1,
		2,
	},
	{
		"empty",
		[][]int{},
		false,
		0,
		0,
	},
	{
		"wrong dimensions",
		[][]int{
			{1, 2},
			{1},
		},
		false,
		0,
		0,
	},
}

var testsMultiply = []struct {
	description string
	in1         [][]int
	in2         [][]int
	ok          bool
	result      [][]int
}{
	{
		"basic2x2",
		[][]int{
			{1, 2},
			{2, 3},
		},
		[][]int{
			{1, 2},
			{2, 3},
		},
		true,
		[][]int{
			{5, 8},
			{8, 13},
		},
	},
	{
		"fail3x2 x 3x2",
		[][]int{
			{1, 2, 3},
			{2, 3, 3},
		},
		[][]int{
			{1, 2, 3},
			{2, 3, 3},
		},
		false,
		nil,
	},
	{
		"diffdimensions3x2 x 2x3",
		[][]int{
			{1, 2},
			{2, 3},
			{3, 4},
		},
		[][]int{
			{1, 2, 3},
			{2, 3, 4},
		},
		true,
		[][]int{
			{5, 8, 11},
			{8, 13, 18},
			{11, 18, 25},
		},
	},
	{
		"simple3x1 x 1x3",
		[][]int{
			{1},
			{2},
			{3},
		},
		[][]int{
			{1, 2, 3},
		},
		true,
		[][]int{
			{1, 2, 3},
			{2, 4, 6},
			{3, 6, 9},
		},
	},
	{
		"nil",
		[][]int{
			{1},
			{2},
			{3},
		},
		nil,
		false,
		nil,
	},
}

func TestNil(t *testing.T) {
	m, err := New(nil)
	if m != nil || err == nil {
		t.Errorf("Results don't match: actual %v expected %v", m, nil)
	}
}

func TestNew(t *testing.T) {
	for _, tc := range testsNew {
		t.Run(fmt.Sprintf("Test: %s", tc.description), func(t *testing.T) {
			m, err := New(tc.in)
			isOk := err == nil
			if err != nil && !tc.ok {
				return
			}
			if err != nil && tc.ok {
				t.Errorf("Expected no error, but got one: %v", err)
				return
			}
			if err == nil && !tc.ok {
				t.Errorf("Expected an error, but got none")
				return
			}
			if isOk != tc.ok || m.cols != tc.cols || m.rows != tc.rows {
				t.Errorf("Results don't match: actual %v expected %v", m, tc)
				return
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	for _, tc := range testsMultiply {
		t.Run(fmt.Sprintf("Test: %s", tc.description), func(t *testing.T) {
			m1, err := New(tc.in1)
			m2, err := New(tc.in2)
			res, err := m1.Multiply(m2)
			isOk := err == nil
			if err != nil && !tc.ok {
				return
			}
			if err != nil && tc.ok {
				t.Errorf("Expected no error, but got one: %v", err)
				return
			}
			if err == nil && !tc.ok {
				t.Errorf("Expected an error, but got none")
				return
			}
			if isOk != tc.ok || !reflect.DeepEqual(res, tc.result) {
				t.Errorf("Results don't match: actual %v expected %v", res, tc)
				return
			}
		})
	}
}
