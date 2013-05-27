package sorting

import (
    "math/rand"
)

func get_random_pivot(right int) int {
    if right < 1 {
        return right
    }
    pivot := rand.Int() % right
    return pivot
}

func quicksort(s []int32) {
    if len(s) > 1 {
        pivot := get_random_pivot(len(s))
        normalized_index := partition(s, pivot)
        quicksort(s[0:normalized_index])
        quicksort(s[normalized_index+1:len(s)])
    } else {
        return
    }
}

func partition(s []int32, index int) int {
    sentinel := s[index]
    max_index := len(s) - 1
    s[index], s[max_index] = s[max_index], s[index]
    stored := 0
    for i:=0;i<max_index;i++ {
        if s[i] < sentinel {
            s[i], s[stored]  = s[stored], s[i]
            stored++
        }
    }
    s[stored], s[max_index] = s[max_index], s[stored]
    return stored
}
