package sorting

import (
    "fmt"
    "testing"
    "io/ioutil"
    "strconv"
    "strings"
)

var FILENAME_TMPL string = "arraysize-10e%v-test%v.dat"

func readFile(fname string) (nums []int32, err error) {
    b, err := ioutil.ReadFile(fname)
    if err != nil { return nil, err }
    lines := strings.Split(string(b), "\n")
    nums = make([]int32, 0, len(lines))

    for _, l := range lines {
        if len(l) == 0 { continue }
        n, err := strconv.Atoi(l)
        
        if err != nil { return nil, err }
        nums = append(nums, int32(n))
    }

    return nums, nil
}

type algo func([]int32)

func benchmark_algorithm(sorting_algo algo, complexity int, b *testing.B) {
    b.StopTimer()
    var integers []int32
    for j := 0; j < b.N; j++ {
        for i:=1; i<4; i++ {
            filename := fmt.Sprintf(FILENAME_TMPL, complexity, i)
            integers, _ = readFile(filename)
            b.StartTimer()
            sorting_algo(integers)
            b.StopTimer()
        }
    }
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
