package sorting

func merge(l []int32, r []int32) []int32 {
    merged := make([]int32, len(l) + len(r))
    for index:=0; index<len(merged); index++ {
        switch {
            case len(l) > 0 && len(r) > 0:
                switch {
                    case l[0] <= r[0]:
                        merged[index] = l[0]
                        l = l[1:]
                    case true:
                        merged[index] = r[0]
                        r = r[1:]
                }
            case len(l) > 0:
                for _, value := range l {
                    merged[index] = value
                    index++
                }
            case len(r) > 0:
                for _, value := range r {
                    merged[index] = value
                    index++
                }
        
        }
    }
    return merged
}

func mergesort(s []int32) {
    mergesort_helper(s)
}

func mergesort_helper(s []int32) []int32 {
    if len(s) < 2 {
        return s
    }
    pivot := len(s) / 2
    left := mergesort_helper(s[:pivot])
    right := mergesort_helper(s[pivot:])
    copy(s, merge(left, right))
    return s
}
