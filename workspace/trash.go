package main

import "fmt"

func exper(cat ...int) {
    fmt.Println(len(cat))
}

func main() {
    kek := []int{5,15,25}
    exper(kek...)   // 3
}

