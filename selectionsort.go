package sorting

func selectionsort(s []int32) {
    array_len := len(s)
    for i:=0; i<array_len; i++ {
        minimum := i
        for j:=i; j<array_len; j++ {
            if s[j] < s[minimum] {
                minimum = j
            }
        }
        s[i], s[minimum] = s[minimum], s[i]
    }
    return
}
