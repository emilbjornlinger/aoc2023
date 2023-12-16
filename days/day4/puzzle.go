package day4

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2023/input"
    "strings"
    "slices"
)

const dayName string = "day4"

type Card struct {
    number int
    wins []Card
}

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    sum := 0

    for _, line := range inputSlice {
        sum += countLine(line)
    }

    fmt.Printf("Output: %v\n", sum)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    for _, line := range inputSlice {
        fmt.Println(line)

        // Create card type with card number and winning cards

        // Parse all cards and extract winning numbers

        // Loop over all cards in new slice

        // Have a recursive function that will handle one card and then handle
        // the cards that that card has won, each time we enter the recursive
        // function a pointer or global variable is incremented, representing 
        // that that card is won
    }

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}

func countLine(line string) int {
    lineCount := 0

    splitColon := strings.Split(line, ":")
    splitBar := strings.Split(splitColon[1], "|")

    winning := strings.Split(splitBar[0], " ")
    numbers := strings.Split(splitBar[1], " ")

    for _, num := range numbers {
        if slices.Contains(winning, num) {
            if num != "" {
                if lineCount == 0 {
                    lineCount = 1
                } else {
                    lineCount = lineCount*2
                }
            }
        }
    }

    return lineCount
}
