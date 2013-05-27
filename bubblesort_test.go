package sorting

import (
    "fmt"
    "testing"
)

func TestBubbleSort_10e1(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
        integers, _ = readFile(filename)
        bubblesort(integers)
        if !isSorted(integers) {
            t.Error("bubblesort fails with an sorted array.")
        }
    }
}

func TestBubbleSort_10e2(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
        integers, _ = readFile(filename)
        bubblesort(integers)
        if !isSorted(integers) {
            t.Error("bubblesort fails with an sorted array.")
        }
    }
}

func TestBubbleSort_10e3(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
        integers, _ = readFile(filename)
        bubblesort(integers)
        if !isSorted(integers) {
            t.Error("bubblesort fails with an sorted array.")
        }
    }
}

func BenchmarkBubbleSort_10e1(b *testing.B) {
    benchmark_algorithm(bubblesort, 1, b)
}

func BenchmarkBubbleSort_10e2(b *testing.B) {
    benchmark_algorithm(bubblesort, 2, b)
}

func BenchmarkBubbleSort_10e3(b *testing.B) {
    benchmark_algorithm(bubblesort, 3, b)
}

func BenchmarkBubbleSort_10e4(b *testing.B) {
    benchmark_algorithm(bubblesort, 4, b)
}

func BenchmarkBubbleSort_10e5(b *testing.B) {
    benchmark_algorithm(bubblesort, 5, b)
}
