package main

import (
  "fmt"
  "time"
)

func main() {
   timer1 := time.NewTimer(2 * time.Second)
   
   // the following code after timer1 will trigger in 2s
   <-timer1.C
   fmt.Println("Timer 1 fired")

   timer2 := time.NewTimer(time.Second)
   go func() {
     <-timer2.C
     fmt.Println("Timer 2 fired")
   }()
   stop2 := timer2.Stop()
   if stop2 {
      fmt.Println("Timer 2 stopped")
   }

   time.Sleep(2 * time.Second)

}
