package sorting

func insertionsort(s []int32) {
    i := 1
    j := 1
    for i=1; i < len(s); i++ {
        for j=i; j >= 1; j-- {
            if s[j] > s[j-1] {
                break;
            } else {
                s[j], s[j-1] = s[j-1], s[j]
            }
        }
    }
    return
}
