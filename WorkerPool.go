package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
   for j := range jobs {
     fmt.Println("worker", id, "started job", j)
     time.Sleep(time.Second)
     fmt.Println("worker", id, "finished job", j)
     results <- j * 2
   }

}

func main() {
   const numJobs = 5
   jobs := make(chan int, numJobs)
   results := make(chan int, numJobs)

   // initiate workers (remained blocked until jobs
   // start comming in
   for w := 1; w <= 3; w++ {
     go worker(w, jobs, results)
   }
   
   // sent 5 jobs
   for j := 1; j <= numJobs; j++ {
     jobs <-j
   }
   close(jobs)

   //loop through the result
  for a :=1; a <= numJobs; a++ {
     <-results
  }
  
}
