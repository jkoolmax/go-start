package main

import (
	"fmt"
	"slices"
)

func main() {

   var s []int
   fmt.Println("uninit:", s, s == nil, len(s) == 0)
   
   s = make([]int, 3)

   s[0] = 3
   s[1] = 4
   s[2] = 5
   fmt.Println("set:", s)
   fmt.Println("get:", s[1])

   fmt.Println("len:", len(s))
   
   s = append(s, 14)
   s = append(s, 8, 9)
   fmt.Println("apd:", s)

   c := make([]int, len(s))
   copy(c, s)
   fmt.Println("cpy:", c)

   l := s[2:5]
   fmt.Println("sl1:", l)

   t := []int{2, 4, 5}
   t2 := []int{2, 4, 5}
   if slices.Equal(t, t2) {
	   fmt.Println("t == t2")
   }
}
