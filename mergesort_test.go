package sorting

import (
    "fmt"
    "testing"
)

func TestMergeSort_merge(t *testing.T) {
    left := []int32 {1}
    right := []int32 {2}
    result := merge(left, right)
    if !isSorted(result) {
        fmt.Println(result)
        t.Error("merge fails with an sorted array.")
    }
}

func TestMergeSort_merge_moreEntries(t *testing.T) {
    left := []int32 {1, 2, 3}
    right := []int32 {2, 5, 6}
    result := merge(left, right)
    if !isSorted(result) {
        fmt.Println(result)
        t.Error("merge fails with an sorted array.")
    }
}

func TestMergeSort_10e1(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
        integers, _ = readFile(filename)
        mergesort(integers)
        if !isSorted(integers) {
            t.Error("mergesort fails with an sorted array.")
        }
    }
}

func TestMergeSort_10e2(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
        integers, _ = readFile(filename)
        mergesort(integers)
        if !isSorted(integers) {
            t.Error("mergesort fails with an sorted array.")
        }
    }
}

func TestMergeSort_10e3(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
        integers, _ = readFile(filename)
        mergesort(integers)
        if !isSorted(integers) {
            t.Error("mergesort fails with an sorted array.")
        }
    }
}

func BenchmarkMergeSort_10e1(b *testing.B) {
    benchmark_algorithm(mergesort, 1, b)
}

func BenchmarkMergeSort_10e2(b *testing.B) {
    benchmark_algorithm(mergesort, 2, b)
}

func BenchmarkMergeSort_10e3(b *testing.B) {
    benchmark_algorithm(mergesort, 3, b)
}

func BenchmarkMergeSort_10e4(b *testing.B) {
    benchmark_algorithm(mergesort, 4, b)
}

func BenchmarkMergeSort_10e5(b *testing.B) {
    benchmark_algorithm(mergesort, 5, b)
}
