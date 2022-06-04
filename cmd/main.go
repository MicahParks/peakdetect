package main

import (
	"fmt"

	"github.com/MicahParks/peakdetect"
)

var (
	inputs  = []float64{1, 1, 1.1, 1, 0.9, 1, 1, 1.1, 1, 0.9, 1, 1.1, 1, 1, 0.9, 1, 1, 1.1, 1, 1, 1, 1, 1.1, 0.9, 1, 1.1, 1, 1, 0.9, 1, 1.1, 1, 1, 1.1, 1, 0.8, 0.9, 1, 1.2, 0.9, 1, 1, 1.1, 1.2, 1, 1.5, 1, 3, 2, 5, 3, 2, 1, 1, 1, 0.9, 1, 1, 3, 2.6, 4, 3, 3.2, 2, 1, 1, 0.8, 4, 4, 2, 2.5, 1, 1, 1}
	outputs = []peakdetect.Signal{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 0}
)

func main() {
	println(fmt.Sprintf("Inputs: %d\nOutputs: %d", len(inputs), len(outputs)))

	detector := peakdetect.NewPeakDetector()
	config := peakdetect.Config{
		Influence: 0,
		Threshold: 5,
	}
	err := detector.Initialize(config, inputs[0:30])
	if err != nil {
		panic(err.Error())
	}

	for i, v := range inputs[30:] {
		signal := detector.Next(v)
		println(fmt.Sprintf("Index: %d\n  Signal: %d\n  Match: %t", i, signal, outputs[i+30] == signal))
	}
}
