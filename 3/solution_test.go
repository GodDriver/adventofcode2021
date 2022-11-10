package main

import (
    "testing"
)

var nums = []string{"00100",
                    "11110",
                    "10110",
                    "10111",
                    "10101",
                    "01111",
                    "00111",
                    "11100",
                    "10000",
                    "11001",
                    "00010",
                    "01010"}

func TestGetGammaEpsilonRate(t *testing.T) {

    got := Multiply(GetGammaEpsilonRate(nums, "most"),
                    GetGammaEpsilonRate(nums, "least"))
    want := 198

    if got != want {
        t.Errorf("Output %d not equal to expected %d", got, want)
    }
}

func TestGetOxygenCO2Rating(t *testing.T) {

    got := Multiply(GetOxygenCO2Rating(nums, "most"),
                    GetOxygenCO2Rating(nums, "least"))
    want := 230

    if got != want {
        t.Errorf("Output %d not equal to expected %d", got, want)
    }
}
