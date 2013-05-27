package sorting

func bubblesort(s []int32) {
    swapped := true
    array_len := len(s)
    i := 1
    for swapped {
        swapped = false
        for i=1; i<array_len; i++ {
            if s[i] < s[i-1] {
                s[i], s[i-1] = s[i-1], s[i]
                swapped = true
            }
        }
    }
    return
}
