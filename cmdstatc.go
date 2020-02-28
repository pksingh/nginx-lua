package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
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

// validate() will true if no error
func (x *XRecord) validate() bool {
	if x.err != nil {
		return false
	}
	return true
}

// Total() gives the total value of values in XRecord.
func (x *XRecord) Total() *XRecord {
	if x.evalTotal || x.err != nil {
		return x
	}
	x.Register.Total = 0
	for _, v := range x.data {
		x.Register.Total += v
	}
	x.evalTotal = true
	return x
}

// Mean() gives the mean (average) value of values in XRecord.
func (x *XRecord) Mean() *XRecord {
	if x.evalMean || x.err != nil {
		return x
	}
	if !x.evalTotal {
		x.Total()
		if x.err != nil {
			return x
		}
	}
	x.Register.Mean = x.Register.Total / float64(x.length)
	x.evalMean = true
	return x
}

// Variance() evaluates variance of the values in XRecord.
func (x *XRecord) Variance() *XRecord {
	if x.evalVariance || x.err != nil {
		return x
	}
	if !x.evalMean {
		x.Mean()
		if x.err != nil {
			return x
		}
	}
	for _, i := range x.data {
		r := x.Register.Mean - i
		x.Register.Variance += r * r
	}
	x.Register.Variance /= float64(x.length)
	x.evalVariance = true
	return x
}

// Max() gives the biggest value in XRecord.
func (x *XRecord) Max() *XRecord {
	if x.evalMax || x.err != nil {
		return x
	}
	x.Register.MaxValue = x.sortedData[x.length-1]
	x.evalMax = true
	return x
}

// MaxWithIndices() evaluates the biggest value and associated indices in XRecord.
func (x *XRecord) MaxWithIndices() *XRecord {
	if x.evalMaxWithIndices || x.err != nil {
		return x
	}

	if !x.evalMax {
		x.Max()
		if x.err != nil {
			return x
		}
	}
	x.Register.MaxIndices = []int{}
	if x.length == 1 {
		x.Register.MaxIndices = append(x.Register.MaxIndices, 0)
		x.evalMaxWithIndices = true
		return x
	}
	for i, v := range x.data {
		if v == x.Register.MaxValue {
			x.Register.MaxIndices = append(x.Register.MaxIndices, i)
		}
	}
	x.evalMaxWithIndices = true
	return x
}

// Min() gives the smallest value in XRecord.
func (x *XRecord) Min() *XRecord {
	if x.evalMin || x.err != nil {
		return x
	}
	x.Register.MinValue = x.sortedData[0]
	x.evalMax = true
	return x
}

// MinWithIndices evaluates the smallest value and associated indices in XRecord.
func (x *XRecord) MinWithIndices() *XRecord {
	if x.evalMinWithIndices || x.err != nil {
		return x
	}

	if !x.evalMin {
		x.Min()
		if x.err != nil {
			return x
		}
	}
	x.Register.MinIndices = []int{}
	if x.length == 1 {
		x.Register.MinIndices = append(x.Register.MinIndices, 0)
		x.evalMinWithIndices = true
		return x
	}
	for i, v := range x.data {
		if v == x.Register.MinValue {
			x.Register.MinIndices = append(x.Register.MinIndices, i)
		}
	}
	x.evalMinWithIndices = true
	return x
}

// Range() evaluates range of the values in XRecord.
func (x *XRecord) Range() *XRecord {
	if x.evalRange || x.err != nil {
		return x
	}

	if !x.evalMax {
		x.Max()
		if x.err != nil {
			return x
		}
	}

	if !x.evalMin {
		x.Min()
		if x.err != nil {
			return x
		}
	}

	x.Register.Range = x.Register.MaxValue - x.Register.MinValue
	x.evalRange = true
	return x
}

// Modes() evaluates the values appearing most often in XRecord.
func (x *XRecord) Modes() *XRecord {
	if x.evalModes || x.err != nil {
		return x
	}
	occurences := make(map[float64]int)
	x.Register.Modes = []float64{}
	for _, i := range x.data {
		occurences[i]++
		if occurences[i] > x.Register.ModeRepeatCount {
			x.Register.ModeRepeatCount = occurences[i]
		}
	}
	if len(occurences) == x.length {
		x.Register.ModeRepeatCount = 0
		x.evalModes = true
		return x
	}
	for i, v := range occurences {
		if v == x.Register.ModeRepeatCount {
			x.Register.Modes = append(x.Register.Modes, i)
		}
	}
	x.evalModes = true
	return x
}

// Median() evaluates median of the values in XRecord.
func (x *XRecord) Median(sorted bool) *XRecord {
	if (!sorted && x.evalMedian) || (sorted && x.evalSortedMedian) || x.err != nil {
		return x
	}
	if x.length == 1 {
		x.Register.SortedMedian = x.data[0]
		x.Register.Median = x.data[0]
		x.evalSortedMedian = true
		x.evalMedian = true
		return x
	}

	if sorted {
		if !x.even {
			x.Register.SortedMedian = x.sortedData[x.middleIndex]
		} else {
			x.Register.SortedMedian = (x.sortedData[x.middleIndex] + x.sortedData[x.middleIndex-1]) / 2
		}
		x.evalSortedMedian = true
	} else {
		if !x.even {
			x.Register.Median = x.data[x.middleIndex]
		} else {
			x.Register.Median = (x.data[x.middleIndex] + x.data[x.middleIndex-1]) / 2
		}
		x.evalMedian = true
	}
	return x
}

// StandardDeviation() evaluates standard deviation of the values in XRecord.
func (x *XRecord) StandardDeviation() *XRecord {
	if x.evalStandardDeviation || x.err != nil {
		return x
	}

	if !x.evalVariance {
		x.Variance()
		if x.err != nil {
			return x
		}
	}

	x.Register.StandardDeviation = math.Sqrt(x.Register.Variance)
	x.evalStandardDeviation = true
	return x
}

// New() gives an instance of XRecord from float64 array.
// If not initialize successfully, then returns nil.
func New(data []float64) *XRecord {
	if len(data) == 0 {
		return nil
	}
	x := &XRecord{
		data:   data,
		length: len(data),
		even:   false,
	}

	x.sortedData = make([]float64, x.length)
	copy(x.sortedData, x.data)
	sort.Float64s(x.sortedData)
	modInt, modFrac := math.Modf(float64(x.length))
	if modFrac == 0.0 {
		x.even = true
	}
	x.middleIndex = int(modInt / 2)
	return x
}

// Failed() gives true if one or more calculations failed.
func (x *XRecord) Failed() bool {
	if x.err != nil {
		return true
	}
	return false
}

// Print() gives a string with the contents of a XRecord
func (x *XRecord) Print() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Data: %v", x.data))
	sb.WriteString(fmt.Sprintf("\nTotal: %f", x.Register.Total))
	sb.WriteString(fmt.Sprintf("\nMean: %f", x.Register.Mean))
	sb.WriteString(fmt.Sprintf("\nMedian: %f", x.Register.Median))
	sb.WriteString(fmt.Sprintf("\nSorted Median: %f", x.Register.SortedMedian))
	sb.WriteString(fmt.Sprintf("\nMax: %f", x.Register.MaxValue))
	sb.WriteString(fmt.Sprintf("\nMax Indices: %v", x.Register.MaxIndices))
	sb.WriteString(fmt.Sprintf("\nMin: %f", x.Register.MinValue))
	sb.WriteString(fmt.Sprintf("\nMin Indices: %v", x.Register.MinIndices))
	sb.WriteString(fmt.Sprintf("\nVariance: %f", x.Register.Variance))
	sb.WriteString(fmt.Sprintf("\nStandard Deviation: %f", x.Register.StandardDeviation))
	sb.WriteString(fmt.Sprintf("\nModes: %v", x.Register.Modes))
	sb.WriteString(fmt.Sprintf("\nMode Repeat Count: %d", x.Register.ModeRepeatCount))

	return sb.String()
}

func main() {
	rdata := []float64{1, 2, 3, 4.5, 5.4, 6, 7}
	calx := New(rdata)
	calx.Total()
	calx.Mean()
	calx.Max()
	calx.Min()
	calx.MaxWithIndices()
	calx.MinWithIndices()
	calx.Range()
	calx.Modes()
	calx.Variance()
	calx.Median(true)
	calx.Median(false)
	calx.StandardDeviation()
	fmt.Fprintf(os.Stdout, "calx : %v\n", calx)
	fmt.Fprintln(os.Stdout)
	fmt.Fprintf(os.Stdout, "Statistics: \n%s\n", calx.Print())
}
