/* 
to close channel when all workers are complete
using x, y := <-workers where x is a current worker
and y a bool indicating if there still more work to do
*/
package main

import "fmt"

func main() {
 jobs := make(chan int, 5)
 done := make(chan bool)

 go func() {
     for {
        j, more := <-jobs
	if more {
           fmt.Println("received job", j)
	} else {
          fmt.Println("received all jobs")
	  done <- true
	  return
	}
     }
 }()

 for j := 1; j <= 3; j++ {
    jobs <- j
    fmt.Println("sent job", j)
 }
 close(jobs)
 fmt.Println("sent all jobs")

 <-done // waiting for the current worker

 _, ok := <-jobs
 fmt.Println("received more jobs:", ok)

}
