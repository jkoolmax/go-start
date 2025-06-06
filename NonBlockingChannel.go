/* 
 using default within select
 to implement non-blocking channel ops
*/
package main

import (
	"fmt"
	"time"
)

func main() {
 messages := make(chan string)
 signals := make(chan bool)
 
 go func() {
   time.Sleep(1 * time.Second)
   signals <- true
 }()

 select {
   case msg := <-messages:
	fmt.Println("received message", msg)
   default:
       fmt.Println("no message received")
 } 
 // create message
 msg := "hi"

 // skip sending non-blocking
 select {
    case messages <-msg:
       fmt.Println("sent message", msg)
    default:
       fmt.Println("no message sent")
}  

 // receiving msg asynchronously
 // Skip signal channel
 select {
    case msg := <-messages:
       fmt.Println("received message", msg)
    case sig := <-signals:
       fmt.Println("received signal ", sig)
    default:
	fmt.Println("no activity")
 }

}
