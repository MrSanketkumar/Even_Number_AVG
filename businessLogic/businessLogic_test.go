package businesslogic

import (
	"errors"
	"testing"

	//businesslogic "GOLANG/businessLogic"
	"Even_Number_AVG/constant"
)
func TestEvenNumber(t *testing.T) {
	tests := []struct {
		name          string
		numbers       []float64
		expectOutput  []float64
		expectedError error
	}{
		{
			name:          "average of even numbers",
			numbers:       []float64{2, 4, 6, 8, 10},
			expectOutput:  []float64{2, 4, 6, 8, 10},
			expectedError: nil,
		},
		{
			name:          "average of single number",
			numbers:       []float64{1, 2, 3, 4, 5},
			expectOutput:  []float64{2, 4},
			expectedError: nil,
		},
		{
			name:          "average of empty slice",
			numbers:       []float64{},
			expectOutput:  []float64{},
			expectedError: errors.New(constant.ErrorNoNumbersProvidedMessage),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evenNumbers, err := Numbers.EvenNumber(test.numbers)
			if err != nil {
				if err == test.expectedError {
					t.Errorf(constant.ErrorMessageFormat, test.expectedError, err)
				}
			}

			if len(evenNumbers) != len(test.expectOutput) {
				t.Errorf(constant.ErrorMessageFormat, test.expectOutput, evenNumbers)
			}
		})
	}
}

func TestEvenCalulateAverage(t *testing.T) {
	tests := []struct {
		name         string
		numbers      []float64
		expectOutput float64
		expectError  error
	}{
		{
			name:         "average of even numbers",
			numbers:      []float64{2, 4, 6, 8, 10},
			expectOutput: 6,
			expectError:  nil,
		},
		{
			name:         "average of single number",
			numbers:      []float64{1, 2, 3, 4, 5},
			expectOutput: 3,
			expectError:  nil,
		},
		{
			name:         "average of empty slice",
			numbers:      []float64{},
			expectOutput: 0,
			expectError:  errors.New(constant.ErrorNoNumbersProvidedMessage),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			avg, err := Numbers.EvenCalulateAverage(test.numbers)

			if err != nil {
				if err == test.expectError {
					t.Errorf(constant.ErrorMessageFormat, test.expectError, err)
				}
			}
			if avg != test.expectOutput {
				t.Errorf(constant.ErrorMessageFormat, test.expectOutput, avg)
			}

		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name         string
		numbers      []float64
		expectOutput float64
	}{
		{
			name:         "sum of even numbers",
			numbers:      []float64{2, 4, 6, 8, 10},
			expectOutput: 30,
		},
		{
			name:         "sum of single number",
			numbers:      []float64{1, 2, 3, 4, 5},
			expectOutput: 15,
		},
		{
			name:         "sum of empty slice",
			numbers:      []float64{},
			expectOutput: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sum := Numbers.Sum(test.numbers)
			if sum != test.expectOutput {
				t.Errorf(constant.ErrorMessageFormat, test.expectOutput, sum)
			}

		})
	}
}
