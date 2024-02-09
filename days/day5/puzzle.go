package day5

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2023/input"
    "strings"
    "strconv"
)

const dayName string = "day5"
const MaxInt = int((^uint(0)) >> 1)

type Range struct {
    start int
    end int
    outputStart int
}

type CatRange struct {
    start int
    end int
}

// 279289190 too high

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    mappings := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}


    // Create and populate ranges for each mapping
    seedToSoil := make([]Range, 0)
    populate(&seedToSoil, "seed-to-soil", inputSlice)

    soilToFertilizer := make([]Range, 0)
    populate(&soilToFertilizer, "soil-to-fertilizer", inputSlice)

    fertilizerToWater := make([]Range, 0)
    populate(&fertilizerToWater, "fertilizer-to-water", inputSlice)

    waterToLight := make([]Range, 0)
    populate(&waterToLight, "water-to-light", inputSlice)

    lightToTemperature := make([]Range, 0)
    populate(&lightToTemperature, "light-to-temperature", inputSlice)

    temperatureToHumidity := make([]Range, 0)
    populate(&temperatureToHumidity, "temperature-to-humidity", inputSlice)

    humidityToLocation := make([]Range, 0)
    populate(&humidityToLocation, "humidity-to-location", inputSlice)

    mappingsToRange := make(map[string][]Range)
    mappingsToRange["seed-to-soil"] = seedToSoil
    mappingsToRange["soil-to-fertilizer"] = soilToFertilizer
    mappingsToRange["fertilizer-to-water"] = fertilizerToWater
    mappingsToRange["water-to-light"] = waterToLight
    mappingsToRange["light-to-temperature"] = lightToTemperature
    mappingsToRange["temperature-to-humidity"] = temperatureToHumidity
    mappingsToRange["humidity-to-location"] = humidityToLocation

    // Extract seed numbers
    seedsString := strings.Split(inputSlice[0], " ")
    seedsString = append([]string(nil), seedsString[1:]...)
    seeds := make([]int, 0)
    for _, seed := range seedsString {
        seedNum, _ := strconv.Atoi(seed)
        seeds = append(seeds, seedNum)
    }

    lowestLocation := MaxInt
    for i, seed := range seeds {
        input := seed
        for _, mapping := range mappings {
            input = transform(mappingsToRange[mapping], input)
        }

        fmt.Printf("output for seed %v: %v\n", i, input)

        if input < lowestLocation {
            lowestLocation = input
        }
    }

    fmt.Printf("Output: %v\n", lowestLocation)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "test.txt")
    inputSlice := input.GetInputSlice(filename)

    mappings := []string{"seed-to-soil", "soil-to-fertilizer", "fertilizer-to-water", "water-to-light", "light-to-temperature", "temperature-to-humidity", "humidity-to-location"}

    // Create and populate ranges for each mapping
    seedToSoil := make([]Range, 0)
    populate(&seedToSoil, "seed-to-soil", inputSlice)

    soilToFertilizer := make([]Range, 0)
    populate(&soilToFertilizer, "soil-to-fertilizer", inputSlice)

    fertilizerToWater := make([]Range, 0)
    populate(&fertilizerToWater, "fertilizer-to-water", inputSlice)

    waterToLight := make([]Range, 0)
    populate(&waterToLight, "water-to-light", inputSlice)

    lightToTemperature := make([]Range, 0)
    populate(&lightToTemperature, "light-to-temperature", inputSlice)

    temperatureToHumidity := make([]Range, 0)
    populate(&temperatureToHumidity, "temperature-to-humidity", inputSlice)

    humidityToLocation := make([]Range, 0)
    populate(&humidityToLocation, "humidity-to-location", inputSlice)

    mappingsToRange := make(map[string][]Range)
    mappingsToRange["seed-to-soil"] = seedToSoil
    mappingsToRange["soil-to-fertilizer"] = soilToFertilizer
    mappingsToRange["fertilizer-to-water"] = fertilizerToWater
    mappingsToRange["water-to-light"] = waterToLight
    mappingsToRange["light-to-temperature"] = lightToTemperature
    mappingsToRange["temperature-to-humidity"] = temperatureToHumidity
    mappingsToRange["humidity-to-location"] = humidityToLocation

    ranges := make([]CatRange, 0)

    // Extract seed numbers into slice of ranges
    extractSeedRanges(&ranges, inputSlice[0])

    for _, r := range ranges {
        r.Print()
    }

    // Get output ranges
    lowestLocation := MaxInt
    for _, mapping := range mappings {
        transformCategory(&ranges, mappingsToRange[mapping])
    }

    // Extract lowest location from ranges
    for _, r := range ranges {
        r.Print()
    }


    fmt.Printf("Output: %v\n", lowestLocation)
}

func populate(rangeList *[]Range, mapping string, input []string) {
    parsing := false

    index := 0
    for true {
        if strings.Contains(input[index], mapping) {
            parsing = true
            index++
            continue
        }

        if parsing {
            if (input[index] == "") {
                break
            }

            numStrings := strings.Split(input[index], " ")
            
            start, _ := strconv.Atoi(numStrings[1])
            rangeLength, _ := strconv.Atoi(numStrings[2])
            end := start + rangeLength - 1
            outputStart, _ := strconv.Atoi(numStrings[0])
            newRange := Range{start: start, end: end, outputStart: outputStart}

            (*rangeList) = append((*rangeList), newRange)
        }

        
        index++
        if index == len(input) {
            break
        }
    }
}

func transform(rangeList []Range, input int) int {
    // Loop through all ranges and see if match
    for _, currRange := range rangeList {
        if currRange.start <= input && input <= currRange.end {
            fmt.Printf("output: %v\n", input)
            return currRange.outputStart + (input - currRange.start)
        }
    }

    // Otherwise return the input
    fmt.Printf("output: %v\n", input)
    return input
}

func extractSeedRanges(rangeList *[]CatRange, input string) {

    seedsString := strings.Split(input, " ")
    seedsString = append([]string(nil), seedsString[1:]...)
    for i := 0; i < len(seedsString)-1; i++ {
        if i % 2 == 0 {
            start, _ := strconv.Atoi(seedsString[i])
            length, _ := strconv.Atoi(seedsString[i+1])
            end := start + length - 1
            newRange := CatRange{start: start, end: end}
            (*rangeList) = append((*rangeList), newRange)
        }
    }
}

/*
 * - Loop through each category range
 * - Loop through mappings 
 * - If start in a mapping range
 *      - Start new catRange and find end of range
 *          - If end of mapping
 *              - Update the start of catRange and tick back index to handle same range
 *              - Set end of newly created range
 *          - Else
 *              - Set end of newly created range
 * - Else
 *      - Start new catRange and find end of range
 *      - Loop through mappings and find smallest start in catRange 
 *          - If start of mapping is in catRange 
 *              - Update the start of catRange and tick back index to handle same range
 *              - Set end of newly created range 
 *          - Else 
 *              - Set end of newly created range
 */
func transformCategory(ranges *[]CatRange, mappings []Range) {
    // Create new temporary catRange that will replace the passed in pointer
    newRanges := make([]CatRange, 0)

    // Loop through each category range
    catIndex := 0
    for catIndex != len(*ranges) {
        for mappingIndex, mapping := range mappings {
            if (*ranges)[catIndex].start < // CONTINUE HERE

        }

        // Check if in a mapping

        catIndex++
    }

    // Return new ranges
    *ranges = newRanges
}

func (r *Range) Print() {
    fmt.Printf("\nstart: %v\nend: %v\noutputStart: %v\n", r.start, r.end, r.outputStart)
}

func (r *CatRange) Print() {
    fmt.Printf("\nstart: %v\nend: %v\n", r.start, r.end)
}
