package sorting

import (
    "testing"
)

func TestInsertionSort_indentityList(t *testing.T) {
    var integers []int32
    sorted := insertionsort(integers)
    if !isSorted(sorted) {
        t.Error("insertionsort fails with an sorted array.")
    }
}
