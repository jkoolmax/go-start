package main

import (
   "fmt"
   "time"
)

func main() {
  p := fmt.Println

  now := time.Now()
  p(now)

  then := time.Date(
     2025, 6, 10, 23, 14, 58, 6513872237, time.UTC)
  
  p(then)

  //year
  p(then.Year())
  p(then.Month())
  p(then.Day())
  p(then.Hour())
  p(then.Minute())
}
