package businesslogic

import (
	
	"testing"
)

func BenchmarkEvenNumber(b *testing.B) {
	numbers := Numbers{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, float64(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numbers.EvenNumber()
	}
}

func BenchmarkCalculateAverage(b *testing.B) {
	numbers := Numbers{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, float64(i))
	}	

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numbers.EvenCalulateAverage()
	}
}		

func BenchmarkSum(b *testing.B) {
	numbers :=Numbers{}
	for i := 0; i < 100; i++ {
		numbers = append(numbers, float64(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numbers.Sum()
	}
}	


