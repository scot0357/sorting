package sorting

import(
    "fmt"
    "testing"
)

func TestHeapSort_max_heapify(t *testing.T) {
    integers := []int32 {2, 4, 3}
    max_heapify(integers, 0, 0)
    if integers[0] != 4 {
        t.Error("root is wrong in heapify")
    }
    if integers[1] != 2 {
        t.Error("left child is wrong in heapify")
    }
    if integers[2] != 3 {
        t.Error("right child is wrong in heapify")
    }
}

func TestHeapSort_do_max_heap(t *testing.T) {
    integers := []int32 {4, 3, 2, 1, 5, 2, 6}
    do_max_heap(integers)
    if integers[0] != 6 {
        t.Error("root is wrong in heapify")
    }
    if integers[1] != 5 {
        t.Error("left child is wrong in heapify")
    }
    if integers[2] != 4 {
        t.Error("right child is wrong in heapify")
    }
}

func TestHeapSort_10e1(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 1, i)
        integers, _ = readFile(filename)
        heapsort(integers)
        if !isSorted(integers) {
            t.Error("heapsort fails with an sorted array.")
        }
    }
}

func TestHeapSort_10e2(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 2, i)
        integers, _ = readFile(filename)
        heapsort(integers)
        if !isSorted(integers) {
            t.Error("heapsort fails with an sorted array.")
        }
    }
}

func TestHeapSort_10e3(t *testing.T) {
    var integers []int32
    for i:=1; i<4; i++ {
        filename := fmt.Sprintf(FILENAME_TMPL, 3, i)
        integers, _ = readFile(filename)
        heapsort(integers)
        if !isSorted(integers) {
            t.Error("heapsort fails with an sorted array.")
        }
    }
}

func BenchmarkHeapSort_10e1(b *testing.B) {
    benchmark_algorithm(heapsort, 1, b)
}

func BenchmarkHeapSort_10e2(b *testing.B) {
    benchmark_algorithm(heapsort, 2, b)
}

func BenchmarkHeapSort_10e3(b *testing.B) {
    benchmark_algorithm(heapsort, 3, b)
}

func BenchmarkHeapSort_10e4(b *testing.B) {
    benchmark_algorithm(heapsort, 4, b)
}

func BenchmarkHeapSort_10e5(b *testing.B) {
    benchmark_algorithm(heapsort, 5, b)
}
