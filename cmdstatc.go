package main

// XRecord facilitate to perform mathematical/statical operations.
type XRecord struct {
	data                  []float64
	sortedData            []float64
	length                int
	even                  bool
	middleIndex           int
	err                   error
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
}

