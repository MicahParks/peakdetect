package main

const (
	SignalNegative = -1
	SignalNone     = 0
	SignalPositive = 1
)

type Signal int8

type peakDetector struct {
	index       uint
	influence   float64
	lag         uint
	meanCache   []float64
	stdDevCache []float64
	threshold   float64
}

type PeakDetector interface {
	Initialize(config Config, values []float64) error
	Next(value float64) Signal
	NextBatch(values []float64) []Signal
}

type peakDetector struct{}

func (p *peakDetector) Initialize(config Config, values []float64) error {
	length := len(values)
	if length == 0 {
		// TODO Return error.
	}
	// TODO Validate config.
	// TODO Remove configuration.
	mean, stdDev := meanStdDev(values)
	p.meanCache = make([]float64, length)
	p.stdDevCache = make([]float64, length)
	p.lag = config.Lag
	p.meanCache[length-1] = mean
	p.stdDevCache[length-1] = stdDev
	return nil
}

func (p peakDetector) Next(value float64) Signal {
	p.cache[p.index] = value
	p.index++
	if p.index == p.lag {
		p.index = 0
	}

	// TODO
}

func (p peakDetector) NextBatch(values []float64) []Signal {
	signals := make([]Signal, len(values))
	for i, v := range values {
		signals[i] = p.Next(v)
	}
	return signals
}

func NewPeakDetector() PeakDetector {
	return peakDetector{}
}

func main() {
	println("compiled & ran")
}
