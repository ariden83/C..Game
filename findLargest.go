package main

import "fmt"
import "os"
import "sort"

// FindLargest returns the largest number in the numbers array
func Solve(weight0 int, weight1 int, weight2 int) int {
    m := []int{weight0, weight1, weight2}
    m_1 := make([]int, len(m))
    copy(m_1, m)
    sort.Ints(m_1 )
    x := numbers[len(m_1) - 1]
    return x
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
