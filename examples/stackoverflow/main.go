package main

import (
	"fmt"
	"log"

	"github.com/MicahParks/peakdetect"
)

// This example is the equivalent of the R example from the algorithm's author.
// https://stackoverflow.com/a/54507329/14797322
func main() {
	data := []float64{1, 1, 1.1, 1, 0.9, 1, 1, 1.1, 1, 0.9, 1, 1.1, 1, 1, 0.9, 1, 1, 1.1, 1, 1, 1, 1, 1.1, 0.9, 1, 1.1, 1, 1, 0.9, 1, 1.1, 1, 1, 1.1, 1, 0.8, 0.9, 1, 1.2, 0.9, 1, 1, 1.1, 1.2, 1, 1.5, 1, 3, 2, 5, 3, 2, 1, 1, 1, 0.9, 1, 1, 3, 2.6, 4, 3, 3.2, 2, 1, 1, 0.8, 4, 4, 2, 2.5, 1, 1, 1}

	// Algorithm configuration from example.
	const (
		lag       = 30
		threshold = 5
		influence = 0
	)

	// Create then initialize the peak detector.
	detector := peakdetect.NewPeakDetector()
	err := detector.Initialize(influence, threshold, data[:lag]) // The length of the initial values is the lag.
	if err != nil {
		log.Fatalf("Failed to initialize peak detector.\nError: %s", err)
	}

	// Start processing new data points and determine what signal, if any they produce.
	//
	// This method, .Next(), is best for when data is being processed in a stream, but this simply iterates over a slice.
	nextDataPoints := data[lag:]
	for i, newPoint := range nextDataPoints {
		signal := detector.Next(newPoint)
		var signalType string
		switch signal {
		case peakdetect.SignalNegative:
			signalType = "negative"
		case peakdetect.SignalNeutral:
			signalType = "neutral"
		case peakdetect.SignalPositive:
			signalType = "positive"
		}

		println(fmt.Sprintf("Data point at index %d has the signal: %s", i+lag, signalType))
	}

	// This method, .NextBatch(), is a helper function for processing many data points at once. It's returned slice
	// should produce the same signal outputs as the loop above.
	signals := detector.NextBatch(nextDataPoints)
	println(fmt.Sprintf("1:1 ratio of batch inputs to signal outputs: %t", len(signals) == len(nextDataPoints)))
}
