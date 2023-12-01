package main

import (
    "os"
    "aoc2023/day1"
    "aoc2023/day2"
    "aoc2023/day3"
    "aoc2023/day4"
    "aoc2023/day5"
    "aoc2023/day6"
    "aoc2023/day7"
    "aoc2023/day8"
    "aoc2023/day9"
    "aoc2023/day10"
    "aoc2023/day11"
    "aoc2023/day12"
    "aoc2023/day13"
    "aoc2023/day14"
    "aoc2023/day15"
    "aoc2023/day16"
    "aoc2023/day17"
    "aoc2023/day18"
    "aoc2023/day19"
    "aoc2023/day20"
    "aoc2023/day21"
    "aoc2023/day22"
    "aoc2023/day23"
    "aoc2023/day24"
    "aoc2023/day25"
)

func main() {

    args := os.Args

    if len(args) < 3 {
        panic("Specify day and part as command line arguments: go run main.go <day> <part>")
    }

    switch args[1] {
    case "1":
        if args[2] == "1" {
            day1.Puzzle1() 
        } else if args[2] == "2" {
            day1.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "2":
        if args[2] == "1" {
            day2.Puzzle1() 
        } else if args[2] == "2" {
            day2.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "3":
        if args[2] == "1" {
            day3.Puzzle1() 
        } else if args[2] == "2" {
            day3.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "4":
        if args[2] == "1" {
            day4.Puzzle1() 
        } else if args[2] == "2" {
            day4.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "5":
        if args[2] == "1" {
            day5.Puzzle1() 
        } else if args[2] == "2" {
            day5.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "6":
        if args[2] == "1" {
            day6.Puzzle1() 
        } else if args[2] == "2" {
            day6.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "7":
        if args[2] == "1" {
            day7.Puzzle1() 
        } else if args[2] == "2" {
            day7.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "8":
        if args[2] == "1" {
            day8.Puzzle1() 
        } else if args[2] == "2" {
            day8.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "9":
        if args[2] == "1" {
            day9.Puzzle1() 
        } else if args[2] == "2" {
            day9.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "10":
        if args[2] == "1" {
            day10.Puzzle1() 
        } else if args[2] == "2" {
            day10.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "11":
        if args[2] == "1" {
            day11.Puzzle1() 
        } else if args[2] == "2" {
            day11.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "12":
        if args[2] == "1" {
            day12.Puzzle1() 
        } else if args[2] == "2" {
            day12.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "13":
        if args[2] == "1" {
            day13.Puzzle1() 
        } else if args[2] == "2" {
            day13.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "14":
        if args[2] == "1" {
            day14.Puzzle1() 
        } else if args[2] == "2" {
            day14.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "15":
        if args[2] == "1" {
            day15.Puzzle1() 
        } else if args[2] == "2" {
            day15.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "16":
        if args[2] == "1" {
            day16.Puzzle1() 
        } else if args[2] == "2" {
            day16.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "17":
        if args[2] == "1" {
            day17.Puzzle1() 
        } else if args[2] == "2" {
            day17.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "18":
        if args[2] == "1" {
            day18.Puzzle1() 
        } else if args[2] == "2" {
            day18.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "19":
        if args[2] == "1" {
            day19.Puzzle1() 
        } else if args[2] == "2" {
            day19.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "20":
        if args[2] == "1" {
            day20.Puzzle1() 
        } else if args[2] == "2" {
            day20.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "21":
        if args[2] == "1" {
            day21.Puzzle1() 
        } else if args[2] == "2" {
            day21.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "22":
        if args[2] == "1" {
            day22.Puzzle1() 
        } else if args[2] == "2" {
            day22.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "23":
        if args[2] == "1" {
            day23.Puzzle1() 
        } else if args[2] == "2" {
            day23.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "24":
        if args[2] == "1" {
            day24.Puzzle1() 
        } else if args[2] == "2" {
            day24.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    case "25":
        if args[2] == "1" {
            day25.Puzzle1() 
        } else if args[2] == "2" {
            day25.Puzzle2() 
        } else {
            panic("Invalid second argument")
        }
    default:
        panic("Invalid day")
    }
}
