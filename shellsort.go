package sorting

//~ import(
    //~ "fmt"
//~ )

var GAPS []int = []int {701, 301, 132, 57, 23, 10, 4, 1}

func shellsort(s []int32) {
    for _, gap := range GAPS {
        for i:=0;i<len(s);i+=gap {
            min_index := i
            for j:=i; j<len(s); j+=gap {
                //~ fmt.Printf("Using gap %d\n", gap)
                if s[j] < s[min_index] {
                    min_index = j
                }
            }
            s[i], s[min_index] = s[min_index], s[i]
        }
    }
}
