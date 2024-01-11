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

    fmt.Printf("%v\n", inputSlice[0])

    cards := make([]int, len(inputSlice))
    for i := range cards {
        cards[i] = 1
    }

    output := evaluateCards(&cards, inputSlice)

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

func countLine2(line string) int {
    lineCount := 0

    splitColon := strings.Split(line, ":")
    splitBar := strings.Split(splitColon[1], "|")

    winning := strings.Split(splitBar[0], " ")
    numbers := strings.Split(splitBar[1], " ")

    for _, num := range numbers {
        if slices.Contains(winning, num) {
            if num != "" {
                lineCount++
            }
        }
    }

    return lineCount
}

func evaluateCards(cards *[]int, lines []string) int {
    currentCardIndex := 0
    count := 0

    for (*cards)[len(*cards)-1] != 0 {
        count++
        processCard(cards, currentCardIndex, lines[currentCardIndex])
        (*cards)[currentCardIndex]--

        if (*cards)[currentCardIndex] == 0 {
            currentCardIndex++
            fmt.Printf("index: %v\n", currentCardIndex)
        }
    }

    return count
}

func processCard(cards *[]int, currentCardIndex int, line string) {
    wins := countLine2(line)

    for i := 1; i <= wins; i++ {
        (*cards)[currentCardIndex+i] = (*cards)[currentCardIndex+i]+1
    }
}
