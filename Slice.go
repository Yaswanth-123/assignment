package main

import "fmt"

func input(x []int, err error) []int {
    if err != nil {
        return x
    }
    var d int
    n, err := fmt.Scanf("%d", &d)
    if n == 1 {
        x = append(x, d)
    }
    return input(x, err)
}

func main() {
    fmt.Println("Enter input:")
    x := input([]int{}, nil)
    fmt.Println("Input:", x)
}