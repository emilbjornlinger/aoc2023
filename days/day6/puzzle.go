package day6

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2023/input"
    "strings"
    "strconv"
)

const dayName string = "day6"

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    // Extract numbers as strings
    timeStrings := strings.Split(inputSlice[0], " ")
    distanceStrings := strings.Split(inputSlice[1], " ")
    timeStrings = append([]string(nil), timeStrings[1:]...)
    distanceStrings = append([]string(nil), distanceStrings[1:]...)

    // Purge empty strings resulting from spaces
    timeStrings = PurgeEmptyStrings(timeStrings)
    distanceStrings = PurgeEmptyStrings(distanceStrings)

    // Extract numbers as integers
    times := ConvertSliceOfStringsToIntegers(timeStrings)
    distances := ConvertSliceOfStringsToIntegers(distanceStrings)

    fmt.Println(times)
    fmt.Println(distances)
    
    combinations := make([]int, len(times))

    for i, raceTime := range times {
        currentRaceCounter := 0
        for holdTime := 0; holdTime <= raceTime; holdTime++ {
            travDist := holdTime*(raceTime - holdTime)

            if travDist > distances[i] {
                currentRaceCounter++
            }
        }

        combinations[i] = currentRaceCounter
    }

    multipliedComb := 1 
    for _, comb := range combinations {
        multipliedComb = multipliedComb * comb
    }
    fmt.Printf("Output: %v\n", multipliedComb)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    // Extract numbers as strings
    timeStrings := strings.Split(inputSlice[0], " ")
    distanceStrings := strings.Split(inputSlice[1], " ")
    timeStrings = append([]string(nil), timeStrings[1:]...)
    distanceStrings = append([]string(nil), distanceStrings[1:]...)

    // Purge empty strings resulting from spaces
    timeStrings = PurgeEmptyStrings(timeStrings)
    distanceStrings = PurgeEmptyStrings(distanceStrings)

    timeString := ConcatenateStrings(timeStrings)
    distanceString := ConcatenateStrings(distanceStrings)

    time, err := strconv.Atoi(timeString)
    if err != nil {
        panic (err)
    }
    distance, err := strconv.Atoi(distanceString)
    if err != nil {
        panic (err)
    }

    currentRaceCounter := 0
    for holdTime := 0; holdTime <= time; holdTime++ {
        travDist := holdTime*(time - holdTime)

        if travDist > distance {
            currentRaceCounter++
        }
    }

    fmt.Printf("Output: %v\n", currentRaceCounter)
}

func PurgeEmptyStrings(dirtySlice []string) []string {
    newSlice := make([]string, 0)

    for _, str := range dirtySlice {
        if str != "" {
            newSlice = append(newSlice, str)
        }
    }

    return newSlice
}

func ConvertSliceOfStringsToIntegers(strSlice []string) []int {
    newSlice := make([]int, 0)
    for i := 0; i < len(strSlice); i++ {
        num, err := strconv.Atoi(strSlice[i])
        if err != nil {
            panic (err)
        }

        newSlice = append(newSlice, num)
    }

    return newSlice
}

func ConcatenateStrings(strSlice []string) string {
    newString := ""
    for _, str := range strSlice {
        newString = newString + str
    }

    return newString
}
