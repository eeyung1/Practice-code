package main

import (
    "fmt"
)

func Slice(a []string, nbrs ...int) []string {
    length := len(a)
    
    if len(nbrs) == 0 {
        return nil
    }
    
    start := nbrs[0]
    if start < 0 {
        start = length + start
    }
    
    end := length
    if len(nbrs) >= 2 {
        end = nbrs[1]
        if end < 0 {
            end = length + end
        }
    }
    
    if start < 0 {
        start = 0
    }
    if end > length {
        end = length
    }
    
    if start >= end {
        return nil
    }
    
    return a[start:end]
}

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", Slice(a, 1))
    fmt.Printf("%#v\n", Slice(a, 2, 4))
    fmt.Printf("%#v\n", Slice(a, -3))
    fmt.Printf("%#v\n", Slice(a, -2, -1))
    fmt.Printf("%#v\n", Slice(a, 2, 0))
}