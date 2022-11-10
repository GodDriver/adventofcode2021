package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

type Command struct {
    Direction string
    Value int
}

func GetPositionByInstructions(commands []Command) (int, int) {
    position := 0
    depth := 0
    for _, command := range commands {
        switch command.Direction {
            case "forward":
                position += command.Value
            case "down":
                depth += command.Value
            case "up":
                depth -= command.Value
        }
    }
    return position, depth
}

func GetPositionByInstructionsNewInterpretation(commands []Command) (int, int) {
    position := 0
    depth := 0
    aim := 0
    for _, command := range commands {
        switch command.Direction {
            case "forward":
                position += command.Value
                depth += aim * command.Value
            case "down":
                aim += command.Value
            case "up":
                aim -= command.Value
        }
    }
    return position, depth
}

func Multiply (a, b int) int {
    return a * b
}

func main() {

    var fileName string

    if len(os.Args) == 2 {
        fileName = os.Args[1]
    } else {
        fileName = "input"
    }

    file, file_err := os.Open(fileName)
    if file_err != nil {
        log.Fatalf("failed to open input file")
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    commands := make([]Command, 0)

    for scanner.Scan() {
        line := scanner.Text()
        var command Command
        var parseint_err error
        command.Direction = line[:strings.Index(line, " ")]
        command.Value, parseint_err = strconv.Atoi(line[strings.LastIndex(line, " ") + 1:])
        if parseint_err != nil {
            log.Fatalf("not integer")
        }
        commands = append(commands, command)
    }

    fmt.Printf("The result of the multiplication is %d\n",
               Multiply(GetPositionByInstructions(commands)))
    fmt.Printf("The result of the multiplication in Part Two is %d\n",
               Multiply(GetPositionByInstructionsNewInterpretation(commands)))
}
