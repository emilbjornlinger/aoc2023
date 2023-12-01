package day14

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2023/input"
)

const dayName string = "day14"

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    // Implementation
    for _, line := range inputSlice {
        fmt.Println(line)
    }

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    // Implementation
    for _, line := range inputSlice {
        fmt.Println(line)
    }

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}
