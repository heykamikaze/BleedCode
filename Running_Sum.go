package main
import "fmt"
func runningSum(nums []int) []int {
    result := []int{}
    
    sum := 0
    for _, v := range nums {
        sum = sum + v
        
        result = append(result, sum)
    }
    
    return result
}
 
func main() {
 arr:= [] int {0, 1, 2, 3}
 fmt.Println(runningSum(arr))
}
//cap buildin func counting the capacity of an array same as strlen in Ceee, len func is also ok its even better i guess? append adds the element to the end like lstaddback or smth similar
//range: iterates over elements no need to use len or cap or whatever; the 1st is index, the 2nd is value
