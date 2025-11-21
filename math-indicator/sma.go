// https://www.tradingview.com/support/solutions/43000696841-simple-moving-average/

package mathindicator

func SimpleMovingAverage(prices []float64, period int) []float64 {
	length := len(prices)
	if length == 0 || period <= 0 {
		return []float64{}
	}

	if period > length {
		period = length
	}

	values := make([]float64, length)
	var sum float64

	for i := 0; i < length; i++ {
		sum += prices[i]
		size := i + 1
		if i >= period {
			sum -= prices[i-period]
			size = period
		}

		values[i] = sum / float64(size)
	}

	return values
}
