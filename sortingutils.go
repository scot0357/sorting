package sorting

func readFile(fname string) (nums []int32, err error) {
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }

    lines := strings.Split(string(b), "\n")
    nums = make([]int32, 0, len(lines))

    for i, l := range lines {
        // Empty line occurs at the end of the file when we use Split.
        if len(l) == 0 { continue }
        n, err := strconv.Atoi(l)
        
        if err != nil { return nil, err }
        nums = append(nums, int32(n))
    }

    return nums, nil
}

func isSorted(t []int32) bool {
    i := 1
    for ; i<len(t); i++ {
        if t[i] < t[i-1] {
            return false
        }
    }
    return true
}
