package gocalc

// SumArrayValues - method to sum values from an array
func SumArrayValues(values ...int) int {
	sum := 0
	for _, num := range values {
		sum += num
	}
	return sum
}
