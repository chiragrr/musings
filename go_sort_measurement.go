package main

import (
	"fmt"
)

func reset_a()[]int{
  a := []int {12, 23, 55, 17, 211, 13}
  return (a[:])
}

func main() {
    a := reset_a()
    fmt.Println(a)

    bubblesort(a[:])
    a = reset_a()
    mergesort(a[:])
}

func bubblesort(a []int){
  len_a := len(a)
  for i := 0; i < len_a; i++{
    for j := i; j < len_a; j++{
      if a[j] < a[i]{
        tmp:= a[i]
        a[i] = a[j]
        a[j] = tmp
      }
    }
  }
  fmt.Println(a)
}
func mergesort(a []int){
  first_half := len(a)/2
  j := first_half
  second_half := len(a)
  var copy_array = make ([] int, len(a))
  var i, count = 0, 0
  fmt.Println(a[0:first_half], a[first_half:second_half])
  for i < first_half && j < second_half {
    if a[i] < a[j]{
      copy_array[count] = a[i]
      count = count + 1
      i = i + 1
    } else {
      copy_array[count] = a[j]
      j = j + 1
      count = count + 1
    }
  }
  for i < first_half {
    copy_array[count] = a[i]
    count = count + 1
    i = i + 1
  }
  for j < second_half {
    copy_array[count] = a[j]
    j = j + 1
    count = count + 1
  }

  fmt.Println(copy_array)
}
