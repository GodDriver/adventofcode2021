package main

import (
    "testing"
)

var commands = []Command{
        Command {
            Direction: "forward",
            Value: 5,
        },
        Command {
            Direction: "down",
            Value: 5,
        },
        Command {
            Direction: "forward",
            Value: 8,
        },
        Command {
            Direction: "up",
            Value: 3,
        },
        Command {
            Direction: "down",
            Value: 8,
        },
        Command {
            Direction: "forward",
            Value: 2,
        },
}

func TestGetPositionByInstructions(t *testing.T) {

    got := Multiply(GetPositionByInstructions(commands))
    want := 150

    if got != want {
        t.Errorf("Output %d not equal to expected %d", got, want)
    }
}

func TestGetPositionByInstructionsNewInterpretation(t *testing.T) {

    got := Multiply(GetPositionByInstructionsNewInterpretation(commands))
    want := 900

    if got != want {
        t.Errorf("Output %d not equal to expected %d", got, want)
    }
}
