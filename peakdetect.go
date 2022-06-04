package peakdetect

import (
	"math"
)

const (
	SignalNegative = -1
	SignalNeutral  = 0
	SignalPositive = 1
)

type Signal int8

type peakDetector struct {
	index       uint
	influence   float64
	lag         uint
	cache       []float64
	meanCache   []float64
	stdDevCache []float64
	threshold   float64
}

type PeakDetector interface {
	Initialize(config Config, values []float64) error
	Next(value float64) Signal
	NextBatch(values []float64) []Signal
}

func (p *peakDetector) Initialize(config Config, initialValues []float64) error {
	length := len(initialValues) // TODO This is the lag?
	if length == 0 {
		// TODO Return error.
	}
	p.lag = uint(length)
	p.influence = config.Influence
	p.threshold = config.Threshold
	// TODO Validate config.
	// TODO Remove configuration.
	mean, stdDev := meanStdDev(initialValues)
	p.cache = make([]float64, length)
	copy(p.cache, initialValues)
	p.meanCache = make([]float64, length)
	p.stdDevCache = make([]float64, length)
	p.meanCache[length-1] = mean
	p.stdDevCache[length-1] = stdDev
	p.index = uint(length) - 1
	return nil
}

func (p *peakDetector) Next(value float64) (signal Signal) {
	prevIndex := p.index
	p.index++
	if p.index == p.lag {
		p.index = 0
	}

	if math.Abs(value-p.meanCache[prevIndex]) > p.threshold*p.stdDevCache[prevIndex] {
		p.cache[p.index] = p.influence*value + (1-p.influence)*p.cache[prevIndex]
		if value > p.meanCache[prevIndex] {
			signal = SignalPositive
		} else {
			signal = SignalNegative
		}
	} else {
		signal = SignalNeutral
		p.cache[p.index] = value
	}

	mean, stdDev := meanStdDev(p.cache)
	p.meanCache[p.index] = mean
	p.stdDevCache[p.index] = stdDev

	return signal
}

func (p *peakDetector) NextBatch(values []float64) []Signal {
	signals := make([]Signal, len(values))
	for i, v := range values {
		signals[i] = p.Next(v)
	}
	return signals
}

func NewPeakDetector() PeakDetector {
	return &peakDetector{}
}

func main() {
	println("compiled & ran")
}
