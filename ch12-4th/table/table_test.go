package table

import "testing"

type myTest struct {
	a        int
	b        int
	resInt   int
	resFloat float64
}

var tests = []myTest{
	{a: 1, b: 2, resInt: 0, resFloat: 0.5},
	{a: 5, b: 10, resInt: 0, resFloat: 0.35},
	{a: 2, b: 2, resInt: 1, resFloat: 1.20},
	{a: 4, b: 2, resInt: 2, resFloat: 2.0},
	{a: 5, b: 2, resInt: 2, resFloat: 2.5},
	{a: 5, b: 4, resInt: 1, resFloat: 1.2},
}

func TestAll(t *testing.T)  {
	// --- FAIL: TestAll (0.00s)
    // table_test.go:32: Expected 0.350000, got 0.500000
    // table_test.go:32: Expected 1.200000, got 1.000000
    // table_test.go:32: Expected 1.200000, got 1.250000
	// FAIL
	// FAIL	cty.sh/table	0.442s
	// FAIL
	
	t.Parallel()
	
	for _, test := range tests {
		intR := intDiv(test.a, test.b)
		if intR != test.resInt {
			t.Errorf("Expected %d, got %d", test.resInt, intR)
		}

		floatR := floatDiv(test.a, test.b)
		if floatR != test.resFloat {
			t.Errorf("Expected %f, got %f", test.resFloat, floatR)
		}
	}
}