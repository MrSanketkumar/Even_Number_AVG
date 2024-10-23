package businesslogic

import (
	"Even_Number_AVG/constant"
	"errors"
)

type Numbers []float64

var (
	ErrNoNumbersProvided = errors.New(constant.ErrorNoNumbersProvidedMessage)
)

func (n Numbers) EvenNumber() (Numbers, error) {
	if len(n) == 0 {
		return nil, ErrNoNumbersProvided
	}

	var evenNumbers Numbers
	for _, number := range n {
		if isEven(number) {
			evenNumbers = append(evenNumbers, number)
		}
		
	}
	return evenNumbers, nil
}

func (n Numbers) EvenCalulateAverage() (float64, error) {
	if len(n) == 0 {
		return 0, ErrNoNumbersProvided
	}

	total := n.Sum()

	return total / float64(len(n)), nil
}

func (n Numbers) Sum() float64 {
	var total float64
	for _, number := range n {
		total += number
	}
	return total
}

func isEven(number float64) bool {
	return number == float64(int(number)) && int(number)%2 == 0
}
