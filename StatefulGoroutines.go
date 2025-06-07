/* 
 Access state by multiple goroutines
 where only 1 goroutine access state at time.
 state/data is owner by another goroutine
*/

package main

import (
     "fmt"
     "math/rand"
     "sync/atomic"
     "time"
)

type readOp struct {
  key int
  resp chan int
}

type writeOp struct {
   key int
   val int
   resp chan bool
}

func main() {
   
    var readOps uint64
    var writeOps uint64

    reads := make(chan readOp)
    writes := make(chan writeOp)
   
    // state owned goroutine
    // await struct from reads and writes
    go func() {
       var state = make(map[int]int)
       for {
          select {
	     case read := <-reads:
		 read.resp <- state[read.key]
	     case write := <-writes:
		 state[write.key] = write.val
		 write.resp <- true

	  }

       }
    }()

    // starting 100 go routines to issue reads
    for range 100 {
        go func() {
           for {
              read := readOp{
		key: rand.Intn(5),
		resp: make(chan int)}
	      reads <- read
	      <-read.resp

	      // add delta | delta is n of atomic job on each goroutine
	      atomic.AddUint64(&readOps, 1)
	      time.Sleep(time.Millisecond)

	   }
	}()
    }
  
  // start 10 writes goroutines
  for range 10 {
    go func() {
        write := writeOp{
           key: rand.Intn(5),
	   val: rand.Intn(100),
	   resp: make(chan bool)}

	  writes <-write
	  <-write.resp
	  atomic.AddUint64(&writeOps, 1)
	  time.Sleep(time.Millisecond)
    }()
  }

  // let goroutines work for a second
  time.Sleep(time.Second)

  readOpsFinal := atomic.LoadUint64(&readOps)
  fmt.Println("readOps:", readOpsFinal)
  writeOpsFinal := atomic.LoadUint64(&writeOps)
  fmt.Println("writeOps:", writeOpsFinal)
}

