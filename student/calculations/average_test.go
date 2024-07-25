package calculations

import (
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name   string
		input  []float64
		expect float64
	}{
		{
			name:   "Test average of positive numbers",
			input:  []float64{1, 2, 3, 4, 5},
			expect: 3,
		},
		{
			name:   "Test average of negative numbers",
			input:  []float64{-1, -2, -3, -4, -5},
			expect: -3,
		},
		{
			name:   "Test average of mixed numbers",
			input:  []float64{-1, 2, -3, 4, -5},
			expect: -0.6,
		},
		{
			name:   "Test average of single number",
			input:  []float64{10},
			expect: 10,
		},
		
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Average(test.input)
			if result != test.expect {
				t.Errorf("Expected %f, but got %f", test.expect, result)
			}
		})
	}
}
