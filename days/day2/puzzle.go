package day2

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2023/input"
    "strings"
    "strconv"
)

const dayName string = "day2"

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    sum := 0
    currID := 0
    for _, line := range inputSlice {
        currID++
        possible := true

        setString := strings.Split(line, ":")[1]
        sets := strings.Split(setString, ";")
        
        for _, countString := range sets {
            counts := strings.Fields(countString)
            
            for i := 0; i < len(counts); i += 2 {
                num, err := strconv.Atoi(counts[i])
                if err != nil {
                    panic(err)
                }

                if strings.Contains(counts[i+1], "red") && num > 12 {
                    possible = false 
                    break
                } else if strings.Contains(counts[i+1], "green") && num > 13 {
                    possible = false
                    break
                } else if strings.Contains(counts[i+1], "blue") && num > 14 {
                    possible = false
                    break
                }
            }
        }

        if possible {
            sum = sum + currID
        }
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

    sum := 0
    for _, line := range inputSlice {
        setString := strings.Split(line, ":")[1]
        sets := strings.Split(setString, ";")

        minRed := 0
        minGreen := 0
        minBlue := 0
        
        for _, countString := range sets {
            counts := strings.Fields(countString)
            
            for i := 0; i < len(counts); i += 2 {
                num, err := strconv.Atoi(counts[i])
                if err != nil {
                    panic(err)
                }

                if strings.Contains(counts[i+1], "red") {
                    if num > minRed {
                        minRed = num
                    }
                } else if strings.Contains(counts[i+1], "green") {
                    if num > minGreen {
                        minGreen = num
                    }
                } else if strings.Contains(counts[i+1], "blue") {
                    if num > minBlue {
                        minBlue = num
                    }
                }
            }
        }

        power := minRed * minGreen * minBlue
        sum += power
    }

    fmt.Printf("Output: %v\n", sum)
}
