package day3

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2023/input"
    "strconv"
    "slices"
)

const dayName string = "day3"

type Pair struct {
    row int
    col int
}

type GearPair struct {
    first int
    second int
}

type Counter struct {
    nextNumStart Pair
    nextNumEnd Pair
    nextCharToParse Pair
}

var (
    startPairs = []Pair{Pair{row: -1, col: 0}, Pair{row: -1, col: -1}, Pair{row: 0, col: -1}, Pair{row: 1, col: -1}, Pair{row: 1, col: 0}}
    normalPairs = []Pair{Pair{row: -1, col: 0}, Pair{row: 1, col: 0}}
    endPairs = []Pair{Pair{row: -1, col: 0}, Pair{row: -1, col: 1}, Pair{row: 0, col: 1}, Pair{row: 1, col: 1}, Pair{row: 1, col: 0}}
    adjSymbols = map[string]string{}
    gears = map[int]GearPair{}
    nonGears = map[int]string{}
)

// Between 518896 and 600000
func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    symbols := make([][]string, 0)

    // Populate symbols with input
    for i, line := range inputSlice {
        symbols = append(symbols, make([]string, 0))
        for j := 0; j < len(line); j++ {
            symbols[i] = append(symbols[i], string(line[j]))
        }
    }

    // Get output
    output := countParts(&symbols)

    fmt.Printf("Output: %v\n", output)

    fmt.Printf("Symbols: \n")
    for key, _ := range adjSymbols {
        fmt.Printf("%v\n", key)
    }
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    symbols := make([][]string, 0)

    // Populate symbols with input
    for i, line := range inputSlice {
        symbols = append(symbols, make([]string, 0))
        for j := 0; j < len(line); j++ {
            symbols[i] = append(symbols[i], string(line[j]))
        }
    }

    // Get output
    output := countGears(&symbols)

    fmt.Printf("Output: %v\n", output)
}

func countParts(input *[][]string) int {
    // Initialize counter
    counter := &Counter{nextNumStart: Pair{row: 0, col: 0},
        nextNumEnd: Pair{row: 0, col: 0},
        nextCharToParse: Pair{row: 0, col: 0},
    }

    // Set sum to 0
    sum := 0

    // Get start and end index of next number
    for setStartAndEnd(input, counter) {
        // Check surrounding of number for symbol
        if nextToSymbol(input, counter) {
            // Convert current number to integer and add to sum
            debug := convertToInt(input, counter)
            //sum += convertToInt(input, counter)
            sum += debug
            //fmt.Println(debug)
            //fmt.Printf("Running sum: %v\n", sum)
        }
    }

    return sum
}

func countGears(input *[][]string) int {
    // Initialize counter
    counter := &Counter{nextNumStart: Pair{row: 0, col: 0},
        nextNumEnd: Pair{row: 0, col: 0},
        nextCharToParse: Pair{row: 0, col: 0},
    }

    // Get start and end index of next number
    for setStartAndEnd(input, counter) {
        // Check surrounding of number for gear symbol
        gearPos := nextToGear(input, counter)
        gearKeysForCurrNum := make([]int, 0)
        for _, pos := range gearPos {
            // Convert current number to integer and add to map of gears
            num := convertToInt(input, counter)

            // Convert to unique number from gearpos
            key := pos.row * 140 + pos.col
            
            if slices.Contains(gearKeysForCurrNum, key) {
                // Don't count double for single digits
                continue
            } else {
                // Else add the key to slice of keys for current number
                gearKeysForCurrNum = append(gearKeysForCurrNum, key)
            }
            
            // First check if in list of non-gears
            _, exists := nonGears[key]
            if exists {
                // Continue with next gear position
                continue
            }

            value, exists := gears[key]
            if exists && value.second != 0 {
                // Adjacent to more than two numbers, not a gear
                nonGears[key] = "1"
            } else if exists {
                tmpPair := GearPair{first: value.first, second: num}
                gears[key] = tmpPair
            } else {
                tmpPair := GearPair{first: num, second: 0}
                gears[key] = tmpPair
            }
        }
    }

    // Loop over gears and sum up ratios
    sum := 0
    for _, value := range gears {
        fmt.Printf("%v, %v\n", value.first, value.second)
        sum += value.first * value.second
    }
    return sum
}

func setStartAndEnd(input *[][]string, counter *Counter) bool {
    foundStart := false
    foundEnd := false 

    for !foundEnd {
        _, err := strconv.Atoi((*input)[counter.nextCharToParse.row][counter.nextCharToParse.col])
        if err == nil && !foundStart {
            foundStart = true
            counter.nextNumStart = counter.nextCharToParse
        } else if err != nil && foundStart {
            foundEnd = true    
            // Set end to current char minus 1
            counter.nextNumEnd = counter.nextCharToParse
            counter.nextNumEnd.col = counter.nextNumEnd.col - 1
        }

        if len((*input)[counter.nextCharToParse.row]) - 1 == counter.nextCharToParse.col {
            // Check if last character to parse
            if len((*input)) - 1 == counter.nextCharToParse.row {
                // We could get an edge case here if last input is a number,
                // but it isn't so the program won't take care of this case
                return false
            }
            
            // Break the number if switching rows
            if foundStart && !foundEnd {
                fmt.Printf("Bam\n")
                foundEnd = true
                // Set end of num to last column of the row
                counter.nextNumEnd = counter.nextCharToParse
                counter.nextNumEnd.col = counter.nextNumEnd.col
            }
            counter.nextCharToParse.col = 0
            counter.nextCharToParse.row = counter.nextCharToParse.row + 1
        } else {
            counter.nextCharToParse.col = counter.nextCharToParse.col + 1
        }
    }

    return true
}

func nextToSymbol(input *[][]string, counter *Counter) bool {
    for i := 0; i + counter.nextNumStart.col <= counter.nextNumEnd.col; i++ {
        currPos := Pair{row: counter.nextNumStart.row, col: counter.nextNumStart.col + i}

        // Special case with single digit, check around entire number
        if i == 0 && currPos.col == counter.nextNumEnd.col {
            for _, toAdd := range startPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] != "." {
                        _, exists := adjSymbols[(*input)[posToCheck.row][posToCheck.col]]
                        if !exists {
                            adjSymbols[(*input)[posToCheck.row][posToCheck.col]] = "1"
                        }
                        return true
                    }
                }
            }
            for _, toAdd := range endPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] != "." {
                        _, exists := adjSymbols[(*input)[posToCheck.row][posToCheck.col]]
                        if !exists {
                            adjSymbols[(*input)[posToCheck.row][posToCheck.col]] = "1"
                        }
                        fmt.Printf("I hate this one: %v, %v, %v\n", currPos.row, currPos.col, (*input)[currPos.row][currPos.col])
                        return true
                    }
                }
            }
        } else if i == 0 {
            for _, toAdd := range startPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] != "." {
                        _, exists := adjSymbols[(*input)[posToCheck.row][posToCheck.col]]
                        if !exists {
                            adjSymbols[(*input)[posToCheck.row][posToCheck.col]] = "1"
                        }
                        return true
                    }
                }
            }
        } else if currPos.col == counter.nextNumEnd.col {
            for _, toAdd := range endPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] != "." {
                        _, exists := adjSymbols[(*input)[posToCheck.row][posToCheck.col]]
                        if !exists {
                            adjSymbols[(*input)[posToCheck.row][posToCheck.col]] = "1"
                        }
                        return true
                    }
                }
            }
        } else {
            for _, toAdd := range normalPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] != "." {
                        _, exists := adjSymbols[(*input)[posToCheck.row][posToCheck.col]]
                        if !exists {
                            adjSymbols[(*input)[posToCheck.row][posToCheck.col]] = "1"
                        }
                        return true
                    }
                }
            }
        }
    }

    return false
}

func nextToGear(input *[][]string, counter *Counter) []Pair {
    retVal := make([]Pair, 0)

    for i := 0; i + counter.nextNumStart.col <= counter.nextNumEnd.col; i++ {
        currPos := Pair{row: counter.nextNumStart.row, col: counter.nextNumStart.col + i}

        // Special case with single digit, check around entire number
        if i == 0 && currPos.col == counter.nextNumEnd.col {
            for _, toAdd := range startPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] == "*" {
                        retVal = append(retVal, Pair{row: posToCheck.row, col: posToCheck.col})
                        fmt.Printf("row: %v, col: %v\n", posToCheck.row, posToCheck.col)
                    }
                }
            }
            for _, toAdd := range endPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] == "*" {
                        retVal = append(retVal, Pair{row: posToCheck.row, col: posToCheck.col})
                        fmt.Printf("row: %v, col: %v\n", posToCheck.row, posToCheck.col)
                    }
                }
            }
        } else if i == 0 {
            for _, toAdd := range startPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] == "*" {
                        retVal = append(retVal, Pair{row: posToCheck.row, col: posToCheck.col})
                        fmt.Printf("row: %v, col: %v\n", posToCheck.row, posToCheck.col)
                    }
                }
            }
        } else if currPos.col == counter.nextNumEnd.col {
            for _, toAdd := range endPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] == "*" {
                        retVal = append(retVal, Pair{row: posToCheck.row, col: posToCheck.col})
                        fmt.Printf("row: %v, col: %v\n", posToCheck.row, posToCheck.col)
                    }
                }
            }
        } else {
            for _, toAdd := range normalPairs {
                posToCheck := toAdd.add(currPos)
                if outside(input, posToCheck) {
                    continue
                } else {
                    if (*input)[posToCheck.row][posToCheck.col] == "*" {
                        retVal = append(retVal, Pair{row: posToCheck.row, col: posToCheck.col})
                        fmt.Printf("row: %v, col: %v\n", posToCheck.row, posToCheck.col)
                    }
                }
            }
        }
    }

    return retVal
}

func convertToInt(input *[][]string, counter *Counter) int {
    numString := convertToString(input, counter.nextNumStart.row, counter.nextNumStart.col, counter.nextNumEnd.col+1)
    value, err := strconv.Atoi(numString)
    if err != nil {
        panic(err)
    }

    return value
}

func convertToString(input *[][]string, row, startCol, endCol int) string {
    retVal := ""

    for i := startCol; i < endCol; i++ {
        retVal = retVal + (*input)[row][i]
    }
    
    return retVal
}

func (first *Pair) add(second Pair) Pair {
    return Pair{row: first.row + second.row, col: first.col + second.col}
}

func (pos *Pair) Print() {
    fmt.Printf("row: %v, col: %v\n", pos.row, pos.col)
}

func outside(input *[][]string, pos Pair) bool {
    if pos.row < 0 || pos.row >= len((*input)) {
        return true
    }

    if pos.col < 0 || pos.col >= len((*input)[pos.row]) {
        return true
    }

    return false
}
