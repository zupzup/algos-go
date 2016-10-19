package longestcommonsubsequence

import (
	"fmt"
	"testing"
)

var testsSort = []struct {
	description string
	in1         []string
	in2         []string
	result      int
	str         string
}{
	{
		"basic",
		[]string{"B", "A", "N", "A", "N", "A"},
		[]string{"A", "T", "A", "N", "A"},
		4,
		"AANA",
	},
	{
		"another",
		[]string{"A", "G", "C", "A", "T"},
		[]string{"G", "A", "C"},
		2,
		"AC",
	},
	{
		"same",
		[]string{"A", "G", "C", "A", "T"},
		[]string{"A", "G", "C", "A", "T"},
		5,
		"AGCAT",
	},
	{
		"different",
		[]string{"M", "Z", "J", "A", "W", "X", "U"},
		[]string{"X", "M", "J", "Y", "A", "U", "Z"},
		4,
		"MJAU",
	},
	{
		"empty",
		[]string{},
		[]string{},
		0,
		"",
	},
	{
		"nil",
		nil,
		nil,
		0,
		"",
	},
}

func TestLCS(t *testing.T) {
	for _, tc := range testsSort {
		t.Run(fmt.Sprintf("Test: %s", tc.description), func(t *testing.T) {
			res, str := LCS(tc.in1, tc.in2)
			if res != tc.result || str != tc.str {
				t.Errorf("Error, actual: %v %v expected: %v %v", res, tc.result, str, tc.str)
				return
			}
		})
	}
}
