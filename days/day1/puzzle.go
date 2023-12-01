package day1

import (
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "aoc2023/input"
)

const dayName string = "day1"

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    input := input.GetInputSlice(filename)

    counter := 0

    for _, currLine := range input {
        nums := make([]int, 0)
        for i := 0; i < len(currLine); i++ {
            if num, err := strconv.Atoi(currLine[i:i+1]); err == nil {
                nums = append(nums, num)
            } 
        }

        firstNumString := strconv.Itoa(nums[0])
        secondNumString := strconv.Itoa(nums[len(nums)-1])
        newNum, err := strconv.Atoi(firstNumString + secondNumString)
        if err != nil {
            panic(err)
        }
        counter += newNum
    }

    fmt.Printf("Output: %v\n", counter)
}

// Should be between 55680 and 55821
func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    input := input.GetInputSlice(filename)

    counter := 0

    for _, currLine := range input {
        nums := make([]int, 0)
        handleLine(currLine, &nums)

        firstNumString := strconv.Itoa(nums[0])
        secondNumString := strconv.Itoa(nums[len(nums)-1])
        newNum, err := strconv.Atoi(firstNumString + secondNumString)
        if err != nil {
            panic(err)
        }
        counter += newNum
    }

    fmt.Printf("Output: %v\n", counter)
}

func handleLine(line string, nums *[]int) {
    textNums := []string{
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    }

    for i := 0; i < len(line); i++ {
        if num, err := strconv.Atoi(line[i:i+1]); err == nil {
            *nums = append(*nums, num)
        } else {
            for j := 0; j < len(textNums); j++ {
                if i + len(textNums[j]) <= len(line) &&
                    line[i:i+len(textNums[j])] == textNums[j] {
                    *nums = append(*nums, j+1)
                    break
                }
            }
        } 
    }
}
