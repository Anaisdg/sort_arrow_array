package arrowsort

import (
	"fmt"

	"github.com/apache/arrow/go/arrow/array"
)

// import (
// 	"fmt"
// )

func Sort(arrow *array.Float64) []int {
	// l := len(slice)

	fmt.Println("New Test Case")
	// matrix := matrix{Indices: slice, Values: a}
	// sort.Sort(matrix)
	// return matrix
	var n = len(arrow.Float64Values())
	usedIndices := make([]int, n)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	for i := 0; i < n; i++ {
		var minIdx = i
		for _, j := range usedIndices {
			if arrow.Value(j) < arrow.Value(minIdx) {
				minIdx = j
			}
		}
		indices[i] = minIdx
		usedIndices = append(usedIndices[:minIdx], usedIndices[minIdx:]...)
		fmt.Printf("Lowest value is at index %v which is %v\n", minIdx, arrow.Value(minIdx))
	}

	return indices
}

// for _, vs := range a.mad {
// 	b := arrow.NewFloatBuilder(nil)
// 	b.Resize(vs.Len())
// 	v := vs.Value(i)
// 		diff := v - median
// 		b.Append(diff)
// 	}

// func sort() {
// 	s := [3]int{3, 7, 4, 5}
// 	l := len(s)
// 	a := make([]int, l)
// 	for i := 0; i < l; i++ {
// 		a[i] = i
// 	}
// 	for i := 0; i < l; i++ {
// 		if s[i] < s[i+1] {
// 			fmt.Println(a[i])
// 		} else {
// 			fmt.Println(a[i+1])
// 		}
// 	}
