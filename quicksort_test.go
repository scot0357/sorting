package sorting

import (
    "fmt"
    "testing"
)

func TestQuickSort_10e1(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
        integers, _ = readFile(filename)
        quicksort(integers)
        if !isSorted(integers) {
            t.Error("quicksort fails with an sorted array.")
        }
    }
}

func TestQuickSort_10e2(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
        integers, _ = readFile(filename)
        quicksort(integers)
        if !isSorted(integers) {
            t.Error("quicksort fails with an sorted array.")
        }
    }
}

func TestQuickSort_10e3(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
        integers, _ = readFile(filename)
        quicksort(integers)
        if !isSorted(integers) {
            t.Error("quicksort fails with an sorted array.")
        }
    }
}

func BenchmarkQuickSort_10e1(b *testing.B) {
    benchmark_algorithm(quicksort, 1, b)
}

func BenchmarkQuickSort_10e2(b *testing.B) {
    benchmark_algorithm(quicksort, 2, b)
}

func BenchmarkQuickSort_10e3(b *testing.B) {
    benchmark_algorithm(quicksort, 3, b)
}

func BenchmarkQuickSort_10e4(b *testing.B) {
    benchmark_algorithm(quicksort, 4, b)
}

func BenchmarkQuickSort_10e5(b *testing.B) {
    benchmark_algorithm(quicksort, 5, b)
}

func TestQuickSort_pivot(t *testing.T) {
    right := 64
    pivot := get_random_pivot(right)
    if pivot < 0 || pivot > right {
        t.Error("Bad pivots")
    }
}

func TestQuickSort_pivotSame(t *testing.T) {
    right := 31
    pivot := get_random_pivot(right)
    if pivot < 0 || pivot > right {
        t.Error("Bad pivots")
    }
}

func TestQuickSort_pivotClose(t *testing.T) {
    right := 32
    pivot := get_random_pivot(right)
    if pivot < 0 || pivot > right {
        t.Error("Bad pivots")
    }
}
