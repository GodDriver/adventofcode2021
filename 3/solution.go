package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func GetGammaEpsilonRate(nums []string, criterion string) int {
    var rate_str string
    for i := 0; i < len(nums[0]); i++ {
        var column string
        for _, num := range nums {
            column += string(num[i])
        }
        if criterion == "most" {
            if strings.Count(column, "1") > strings.Count(column, "0") {
                rate_str += "1"
            } else {
                rate_str += "0"
            }
        } else if criterion == "least" {
            if strings.Count(column, "1") > strings.Count(column, "0") {
                rate_str += "0"
            } else {
                rate_str += "1"
            }
        } else {
            log.Fatalf("unknown criterion")
        }
    }
    rate, err := strconv.ParseInt(rate_str, 2, 32)
    if err != nil {
        log.Fatalf("not base 2 string representation")
    }
    return int(rate)
}

func GetOxygenCO2Rating(numbers []string, criterion string) int {
    nums := make([]string, len(numbers))
    copy(nums, numbers)
    var value rune
    for i := 0; i < len(nums[0]); i++ {
        var column string
        for _, num := range nums {
            column += string(num[i])
        }
        if criterion == "most" {
            if strings.Count(column, "1") >= strings.Count(column, "0") {
                value = '1'
            } else {
                value = '0'
            }
        } else if criterion == "least" {
            if strings.Count(column, "0") <= strings.Count(column, "1") {
                value = '0'
            } else {
                value = '1'
            }
        } else {
            log.Fatalf("unknown criterion")
        }
        for j := len(nums) - 1; j >= 0; j-- {
            if rune(nums[j][i]) != value && len(nums) > 1 {
                nums[j] = nums[len(nums) - 1]
                nums = nums[:len(nums) - 1]
            }
        }
    }
    rating, err := strconv.ParseInt(nums[0], 2, 32)
    if err != nil {
        log.Fatalf("not base 2 string representation")
    }
    return int(rating)
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

    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    fmt.Printf("The power consumption of the submarine is %d\n",
               Multiply(GetGammaEpsilonRate(lines, "most"),
                        GetGammaEpsilonRate(lines, "least")))
    fmt.Printf("The life support rating of the submarine is %d\n",
               Multiply(GetOxygenCO2Rating(lines, "most"),
                        GetOxygenCO2Rating(lines, "least")))
}
