package sorting

import (
    "fmt"
    "testing"
)

func TestShellSort_10e1(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
        integers, _ = readFile(filename)
        shellsort(integers)
        if !isSorted(integers) {
            t.Error("shellsort fails with an sorted array.")
        }
    }
}

func TestShellSort_10e2(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
        integers, _ = readFile(filename)
        shellsort(integers)
        if !isSorted(integers) {
            t.Error("shellsort fails with an sorted array.")
        }
    }
}

func TestShellSort_10e3(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
        integers, _ = readFile(filename)
        shellsort(integers)
        if !isSorted(integers) {
            t.Error("shellsort fails with an sorted array.")
        }
    }
}

func BenchmarkShellSort_10e1(b *testing.B) {
    benchmark_algorithm(shellsort, 1, b)
}

func BenchmarkShellSort_10e2(b *testing.B) {
    benchmark_algorithm(shellsort, 2, b)
}

func BenchmarkShellSort_10e3(b *testing.B) {
    benchmark_algorithm(shellsort, 3, b)
}

func BenchmarkShellSort_10e4(b *testing.B) {
    benchmark_algorithm(shellsort, 4, b)
}

func BenchmarkShellSort_10e5(b *testing.B) {
    benchmark_algorithm(shellsort, 5, b)
}
