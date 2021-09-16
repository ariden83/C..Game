package main

import "fmt"
import "os"
import "sort"

func Solve(weight0 int, weight1 int, weight2 int) int {
    m := []int{weight0, weight1, weight2}
    m_1 := make([]int, len(m))
    copy(m_1, m)
    sort.Ints(m_1 )
    x :=  Find(m, m_1[2])
    return x
}

func Find(a []int, x int) int {
    for i, n := range a {
        if x == n && n >= 10 && n <= 200 {
            return i
        }
    }
    return 0
}

/* Ignore and do not change the code below */
func main() {
    for {
        var weight0 int
        fmt.Scan(&weight0)
        
        var weight1 int
        fmt.Scan(&weight1)
        
        var weight2 int
        fmt.Scan(&weight2)
        
        stdoutWriter := os.Stdout
        os.Stdout = os.Stderr
        action := Solve(weight0, weight1, weight2)
        os.Stdout = stdoutWriter
        fmt.Println(action)
    }
}
