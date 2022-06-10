package peakdetect_test

import (
	"errors"
	"testing"

	"github.com/MicahParks/peakdetect"
)

const (
	// These configurations are from the R example of the algorithm's author.
	// https://stackoverflow.com/a/54507329/14797322
	exampleLag       = 30
	exampleInfluence = 0
	exampleThreshold = 5

	logFmt = "%s\nError: %s"
)

var (
	// These inputs and outputs are from the R example of the algorithm's author.
	// https://stackoverflow.com/a/54507329/14797322
	exampleInputs  = []float64{1, 1, 1.1, 1, 0.9, 1, 1, 1.1, 1, 0.9, 1, 1.1, 1, 1, 0.9, 1, 1, 1.1, 1, 1, 1, 1, 1.1, 0.9, 1, 1.1, 1, 1, 0.9, 1, 1.1, 1, 1, 1.1, 1, 0.8, 0.9, 1, 1.2, 0.9, 1, 1, 1.1, 1.2, 1, 1.5, 1, 3, 2, 5, 3, 2, 1, 1, 1, 0.9, 1, 1, 3, 2.6, 4, 3, 3.2, 2, 1, 1, 0.8, 4, 4, 2, 2.5, 1, 1, 1}
	exampleOutputs = []peakdetect.Signal{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0}
)

func TestPeakDetector_Initialize(t *testing.T) {
	detector := peakdetect.NewPeakDetector()
	err := detector.Initialize(0, 0, nil)
	if !errors.Is(err, peakdetect.ErrInvalidInitialValues) {
		t.Fatalf("Invalid initilization did not produce error.\n  Expected: %s\n  Actual: %s", peakdetect.ErrInvalidInitialValues, err)
	}
}

func TestPeakDetector_Next(t *testing.T) {
	detector := peakdetect.NewPeakDetector()
	err := detector.Initialize(exampleInfluence, exampleThreshold, exampleInputs[0:exampleLag])
	if err != nil {
		t.Fatalf(logFmt, "Error during initilization.", err)
	}

	for i, v := range exampleInputs[exampleLag:] {
		signal := detector.Next(v)
		exampleSignal := exampleOutputs[i+exampleLag]
		match := exampleSignal == signal
		if !match {
			t.Fatalf("Example signal did not match actual signal.\n  Example: %d\n  Actual: %d", exampleSignal, signal)
		}
	}
}

func TestPeakDetector_NextBatch(t *testing.T) {
	detector := peakdetect.NewPeakDetector()
	err := detector.Initialize(exampleInfluence, exampleThreshold, exampleInputs[0:exampleLag])
	if err != nil {
		t.Fatalf(logFmt, "Error during initilization.", err)
	}

	signals := detector.NextBatch(exampleInputs[exampleLag:])
	for i, signal := range signals {
		exampleSignal := exampleOutputs[i+exampleLag]
		if signal != exampleSignal {
			t.Fatalf("Example signal did not match actual signal.\n  Example: %d\n  Actual: %d", exampleSignal, signal)
		}
	}
}

func BenchmarkPeakDetector_NextBatch(b *testing.B) {
	detector := peakdetect.NewPeakDetector()
	err := detector.Initialize(exampleInfluence, exampleThreshold, exampleInputs[0:exampleLag])
	if err != nil {
		b.Fatalf(logFmt, "Error during initilization.", err)
	}

	detector.NextBatch(exampleInputs[exampleLag:])
}

func TestPeakDetector_SignalNegative(t *testing.T) {
	data := []float64{0, 1, 0, -1, 0, -500}
	const lag = 5

	detector := peakdetect.NewPeakDetector()
	err := detector.Initialize(exampleInfluence, exampleThreshold, data[:lag])
	if err != nil {
		t.Fatalf(logFmt, "Error during initilization.", err)
	}

	signal := detector.Next(data[lag])
	if signal != peakdetect.SignalNegative {
		t.Fatalf("Signal should have been negative.\n  Actual: %d", signal)
	}
}
