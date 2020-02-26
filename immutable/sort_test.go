package arrowsort

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

func TestSort(t *testing.T) {
	// Make a slice of ints with 2 values
	// []int{0, 1}

	
	pool := memory.NewGoAllocator()

	t.Run("ordered", func(t   *testing.T) {
		b := array.NewFloat64Builder(pool)
		defer b.Release()
		b.AppendValues(
			[]float64{3},
			[]bool{true},
		)
		arr := b.NewFloat64Array()
		defer arr.Release()

		got := Sort(arr)
		want := []int{0}

		if !cmp.Equal(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})
	t.Run("ordered", func(t   *testing.T) {
		b := array.NewFloat64Builder(pool)
		defer b.Release()
		b.AppendValues(
			[]float64{4, 1, 16, 12},
			[]bool{true, true, true, true},
		)
		arr := b.NewFloat64Array()
		defer arr.Release()

		got := Sort(arr)
		want := []int{1,0,3,2}

		
		if !cmp.Equal(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})

}
