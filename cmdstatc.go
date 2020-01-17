package main

import (
	"fmt"
	"os"
)

// XRecord facilitate to perform mathematical/statical operations.
type XRecord struct {
	data                  []float64
	sortedData            []float64
	length                int
	even                  bool
	middleIndex           int
	err                   error // any elaluation error will be strored here
	evalTotal             bool
	evalMean              bool
	evalMedian            bool
	evalRange             bool
	evalVariance          bool
	evalStandardDeviation bool
	evalSortedMedian      bool
	evalMax               bool
	evalMaxWithIndices    bool
	evalMin               bool
	evalMinWithIndices    bool
	evalModes             bool
	Register              Register // for storing evaluated results
}

// Register a way to store evaluated results.
type Register struct {
	Total             float64
	Mean              float64
	Median            float64
	Range             float64
	Variance          float64
	StandardDeviation float64
	SortedMedian      float64
	MaxIndices        []int
	MaxValue          float64
	MinIndices        []int
	MinValue          float64
	ModeRepeatCount   int
	Modes             []float64
}

// Length() gives XRecord data element size.
func (x *XRecord) Length() int {
	return x.length
}

// Even() gives true if the number of element in the XRecord is even.
func (x *XRecord) Even() bool {
	return x.even
}

func main() {
	rdata := []float64{1, 2, 3, 4.5, 5.4, 6, 7}
	xrec := XRecord{data: rdata}
	fmt.Fprintf(os.Stdout, "Statistics: \nxrec : %v\n", xrec)



}
