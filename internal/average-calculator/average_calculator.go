package average_calculator

type AverageCalculator struct {
}

func New() AverageCalculator {
	return AverageCalculator{}
}

// if numbers len=0 => panic!
// if numbers = nil => panic!
func (ac AverageCalculator) CalculateAverage(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return sum / len(numbers)
}