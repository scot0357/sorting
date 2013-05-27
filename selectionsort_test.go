package sorting

import (
    "fmt"
    "testing"
)

func TestSelectionSort_10e1(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
        integers, _ = readFile(filename)
        selectionsort(integers)
        if !isSorted(integers) {
            t.Error("selectionsort fails with an sorted array.")
        }
    }
}

func TestSelectionSort_10e2(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
        integers, _ = readFile(filename)
        selectionsort(integers)
        if !isSorted(integers) {
            t.Error("selectionsort fails with an sorted array.")
        }
    }
}

func TestSelectionSort_10e3(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
        integers, _ = readFile(filename)
        selectionsort(integers)
        if !isSorted(integers) {
            t.Error("selectionsort fails with an sorted array.")
        }
    }
}

func BenchmarkSelectionSort_10e1(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            selectionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkSelectionSort_10e2(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            selectionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkSelectionSort_10e3(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            selectionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkSelectionSort_10e4(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 4, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            selectionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkSelectionSort_10e5(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 5, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            selectionsort(integers)
            b.StopTimer()
        }
    }
}
