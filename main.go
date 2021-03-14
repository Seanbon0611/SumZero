package main

import "fmt"
/*
write a function called sumZero that accepts a sorted array on integers, and return the integers whos sum is equal to zero
*/
func sumZero(arr []int) ([]int, error){
  start := 0
  end := len(arr) - 1
  var result []int
  for start < end {
    sum := arr[start] + arr[end]
    if sum == 0 {
      result := []int{arr[start], arr[end]}
        fmt.Println(result)
      return result, nil
    } else if sum > 0 {
      end--
    } else {
      start++
    }
  }
  fmt.Println("undefined")
  return result, nil
}

func main() {
  sumZero([]int{-3,-1,0,2,3,5})
  sumZero([]int{-3,-1,0,2,2,5})
}