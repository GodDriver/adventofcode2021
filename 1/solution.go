package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func GetNumOfIncreases(nums []int) int {
    count := 0
    var prev_num int
    for i, num := range nums {
        if i > 0 && num > prev_num {
            count++
        }
	prev_num = num
    }
    return count
}

func GetNumOfSumIncreases(nums []int, num_count int) int {
    count := 0
    var prev_sum int
    sums := GetSums(nums, num_count)
    for i, sum := range sums {
        if i > 0 && sum > prev_sum {
            count++
        }
	prev_sum = sum
    }
    return count
}

func GetSums(nums []int, num_count int) []int {
    var sums []int
    for i, _ := range nums {
        if i >= num_count - 1 {
            var sum int
            for n := 0; n < num_count; n++ {
                sum += nums[i - n]
            }
            sums = append(sums, sum)
        }
    }
    return sums
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

    var depths []int

    for scanner.Scan() {
        depth, parseint_err := strconv.Atoi(scanner.Text())
        if parseint_err != nil {
            log.Fatalf("not integer")
        }
        depths = append(depths, depth)
    }

    fmt.Printf("%d measurements are larger than the previous measurement\n",
               GetNumOfIncreases(depths))
    fmt.Printf("%d sums are larger than the previous sum\n",
               GetNumOfSumIncreases(depths, 3))
}
