package main

import (
  "fmt"
  "math/rand/v2"
)

func main() {
  fmt.Print(rand.IntN(100), ","); fmt.Print(rand.IntN(100))
  fmt.Println()

  fmt.Printf("%.2f\n", rand.Float64())
  
  fmt.Print((rand.Float64() * 5) + 5, ",")
  fmt.Print((rand.Float64() * 5) + 5)
  fmt.Println()

  s2 := rand.NewPCG(42, 1024)
  r2 := rand.New(s2)
  fmt.Print(r2.IntN(100), ",")
  fmt.Print(r2.IntN(100))
  fmt.Println()

}
