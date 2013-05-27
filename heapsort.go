package sorting

func left_child(index int) int {
    return (2 * index) + 1
}

func right_child(index int) int {
    return (2 * index) + 2
}

func do_max_heap(s []int32) {
    i := (len(s) / 2) + 1
    for ;i>=0;i-- {
        max_heapify(s, i, 0)
    }
}

func max_heapify(s []int32, index int, size_mod int) {
    var max int
    left := left_child(index)
    right := right_child(index)
    if left < len(s) - size_mod && s[index] < s[left] {
        max = left
    } else {
        max = index
    }
    if right < len(s) - size_mod && s[max] < s[right] {
        max = right
    }
    if max != index {
        s[max], s[index] = s[index], s[max]
        max_heapify(s, max, size_mod)
    }
}

func heapsort(s []int32) {
    do_max_heap(s)
    i := len(s) - 1
    j := 1
    for ;i>0;i-- {
        s[0], s[i] = s[i], s[0]
        max_heapify(s, 0, j)
        j++
    }
}

