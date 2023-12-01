package input

import (
    "os"
    "bufio"
)

func GetInputSlice(filename string) []string {

    f, err := os.Open(filename)
    if err != nil {
        panic (err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    input := make([]string, 0)

    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    if errScan := scanner.Err(); errScan != nil {
        panic(errScan)
    }

    return input
}
