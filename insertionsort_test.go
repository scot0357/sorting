package sorting

import (
    "fmt"
    "testing"
)

func TestInsertionSort_10e1(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
        integers, _ = readFile(filename)
        insertionsort(integers)
        if !isSorted(integers) {
            t.Error("insertionsort fails with an sorted array.")
        }
    }
}

func TestInsertionSort_10e2(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
        integers, _ = readFile(filename)
        insertionsort(integers)
        if !isSorted(integers) {
            t.Error("insertionsort fails with an sorted array.")
        }
    }
}

func TestInsertionSort_10e3(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
        integers, _ = readFile(filename)
        insertionsort(integers)
        if !isSorted(integers) {
            t.Error("insertionsort fails with an sorted array.")
        }
    }
}

func BenchmarkInsertionSort_10e1(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            insertionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkInsertionSort_10e2(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            insertionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkInsertionSort_10e3(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            insertionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkInsertionSort_10e4(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 4, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            insertionsort(integers)
            b.StopTimer()
        }
    }
}

func BenchmarkInsertionSort_10e5(b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, 5, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            insertionsort(integers)
            b.StopTimer()
        }
    }
}
