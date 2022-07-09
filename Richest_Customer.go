package main

import "fmt"

type account [][]int
func runningSum(nums []int) int {

    sum := 0
    for _, v := range nums {
        sum = sum + v
    }

    return sum
}

func maximumWealth(accounts [][]int) int {
    max := 0
   /* leng := len(accounts)*/
    for _, account := range accounts {
        sum := runningSum(account)
        if sum > max {
            max = sum}
    } 
    return max
}

func main() {
check := account{
    []int{1, 3, 8, 23, 3, 123},
    []int{1, 0, 4},
}
    fmt.Println(maximumWealth(check))
}
