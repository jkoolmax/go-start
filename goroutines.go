package main
import (
	"fmt"
	"time" // imported to sleep a thread
)

func f(from string) {
  for i := range int(len(from)) {
    fmt.Println(from, ":", i)
  }
}

func main() {
  f("direct")

  go f("goroutine")

  go func(msg string){
    fmt.Println(msg)
  }("going")

  time.Sleep(time.Second)
  fmt.Println("done")
}
