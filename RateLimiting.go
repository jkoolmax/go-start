/* 
Request RateLimit simulation is possible using
1. time.Ticker
2. a channel that send and receive time.Time

when time.Ticker is running every x seconds,
if we fill a channel that receives time.Time,
we will trigger a burst of up to 3 events
*/

package main

import (
	"fmt"
	"time"
)

func main() {
     requests := make(chan int, 5)
     for i := 1; i <= 5; i ++ {
         requests <- i
     }

     // after sending tasks to a buffered channel
     // we close it
     close(requests)

     // creating ticker interval that runs every 200ms
     limiter := time.Tick(200 * time.Millisecond)

     // serve each request every 200ms
     for req := range requests {
         <-limiter
	 fmt.Println("request", req, time.Now())
     }

     // create a burstLimit channel on rate of 3
     // it can sent or receive 3 time objects
     burstyLimiter := make(chan time.Time, 3)

     // filling up the channel
     for range 3 {
        burstyLimiter <-time.Now()
     }
   
     // send time obj every 200ms
     go func() {
        for t := range time.Tick(200 * time.Millisecond){
           burstyLimiter <-t
	}
     }()
     
     // simulate 5 more incoming requests
     // the first 3 will benefit from the burstyLimit capability
     burstyRequests := make(chan int, 5)
     for i := 1; i <= 5; i++ {
        burstyRequests <- i
     }
     close(burstyRequests)
     
     for req := range burstyRequests {
        <-burstyLimiter
	fmt.Println("request", req, time.Now())
     }

}
