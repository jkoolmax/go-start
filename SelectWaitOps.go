package main

import (
	"fmt"
	"time"
)

func receiveMsg(msg string) {
   fmt.Printf("received: %s\n", msg)
}
func main() {
  c1 := make(chan string)
  c2 := make(chan string)

  go func() {
    time.Sleep(1 * time.Second)
    c1 <- "one"
  }()

  go func(){
     time.Sleep(2 * time.Second)
     c2 <- "two"
  }()

  for range 2 {
      select {
        case msg1 := <-c1:
	    receiveMsg(msg1)
      case msg2 := <-c2:
	   receiveMsg(msg2)
     }
  }
}
