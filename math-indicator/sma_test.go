package mathindicator

import (
	"fmt"
	"testing"
)

func TestSimpleMovingAverage(t *testing.T) {
	input1 := []float64{10.0, 20.0, 30.0, 40.0, 50.0, 60.0}
	period1 := 3
	result1 := SimpleMovingAverage(input1, period1)
	fmt.Println(result1)

	input2 := []float64{100.0, 102.0, 101.0, 103.0, 105.0}
	period2 := 2
	result2 := SimpleMovingAverage(input2, period2)
	fmt.Println(result2)

	input3 := []float64{5.0, 10.0, 15.0}
	period3 := 4
	result3 := SimpleMovingAverage(input3, period3)
	fmt.Println(result3)
}
