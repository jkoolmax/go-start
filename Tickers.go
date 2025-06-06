package main

import (
	"fmt"
	"time"
)

func main() {
  ticker := time.NewTicker(500 * time.Millisecond)
  done := make(chan bool)


  go func() {
     for {
        select {
          case <-done:
	    return
	  // run every 500ms
    	  case t := <-ticker.C:
		  fmt.Println("Tick at", t)
	}
     }
  }()
  
  // emulate task duration completion of 2.5s 
  time.Sleep(2600 * time.Millisecond)
  
  // stop ticker after 2.6s
  ticker.Stop()
  
  done <-true
  fmt.Println("Ticker stopped")

}
