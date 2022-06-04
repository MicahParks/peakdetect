package peakdetect

import (
	"math"
)

func meanStdDev(sample []float64) (mean, stdDev float64) {
	for _, num := range sample {
		mean += num
	}
	mean /= float64(len(sample))

	for _, num := range sample {
		stdDev += math.Pow(num-mean, 2)
	}
	stdDev /= float64(len(sample))
	stdDev = math.Sqrt(stdDev)

	return mean, stdDev
}
