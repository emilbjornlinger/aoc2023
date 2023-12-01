# aoc2023
Advent of Code 2023

---

There is one binary for all days that is run with 

```
go run main.go <day> <part>
```

from the root directory. E.g.

```
go run main.go 1 2
```

to run part two of the first day.

The command line arguments specify which day and which part of the puzzle that will be executed. Trying to run days where the input file has not been downloaded yet will result in a panic. This can be ammended by downloading the input for that day and placing it in the directory for that day, naming it "input.txt".

---
### Known bugs

There is currently a bug where the filepath for the second puzzle for some days is not generated correctly.
