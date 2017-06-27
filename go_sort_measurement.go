package main

import (
	"fmt"
  "math/rand"
  "time"
)

func reset_a()[]int{
  var a [5000]int
  for i:=0; i < 5000; i++ {
    a[i] =  rand.Intn(65000)
  }
  return (a[:])
}

func main() {
    a := reset_a()
    pre_bub := time.Now().UnixNano()
    bubblesort(a[:])
    post_bub := time.Now().UnixNano()

    a = reset_a()
    pre_merge := time.Now().UnixNano()
    mergesort(a[:])
    post_merge := time.Now().UnixNano()

    fmt.Println("Time taken(nano sec) while sorting 5K Integers using bubble sort :", post_bub - pre_bub)
    fmt.Println("Time taken(nano sec) while sorting 5K Integers using merge sort :", post_merge - pre_merge)
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
  //fmt.Println(a)
}


func mergesort(a []int){
  first_half := len(a)/2
  j := first_half
  second_half := len(a)
  var copy_array = make ([] int, len(a))
  var i, count = 0, 0
  //fmt.Println(a[0:first_half], a[first_half:second_half])
  for i < first_half && j < second_half {
    if a[i] < a[j]{
      copy_array[count] = a[i]
      count = count + 1
      i = i + 1
    } else{
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

  //fmt.Println(copy_array)
}
