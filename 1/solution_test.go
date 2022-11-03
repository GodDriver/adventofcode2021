package solution

import (
    "testing"
)

var nums = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestGetNumOfIncreases(t *testing.T) {

    got := GetNumOfIncreases(nums)
    want := 7

    if got != want {
        t.Errorf("Output %d not equal to expected %d", got, want)
    }
}

func TestGetNumOfSumIncreases(t *testing.T) {

    got := GetNumOfSumIncreases(nums, 3)
    want := 5

    if got != want {
        t.Errorf("Output %d not equal to expected %d", got, want)
    }
}
