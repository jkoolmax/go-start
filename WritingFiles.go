package main

import (
    "bufio"
    "fmt"
    "os"
)

func checkError(e error) {
     if e != nil {
        panic(e)
     }
}

func main() {
	
  d1 := []byte("hello\ngo\n") // dump a byte or string
  err := os.WriteFile("/tmp/text.txt", d1, 0644) //write with perms
  checkError(err)

  f, err := os.Create("/tmp/text2.txt") // doing granular write
  checkError(err)

  defer f.Close() // defer close

  // start writing
  d2 := []byte{115, 111, 109, 101, 10}
  n2, err := f.Write(d2)
  checkError(err)
  fmt.Printf("wrote %d bytes\n", n2)

  n3, err := f.WriteString("writes\n")
  checkError(err)
  fmt.Printf("wrote %d bytes\n", n3)
  
  f.Sync() // flush write to stable storage

  w := bufio.NewWriter(f)
  n4, err := w.WriteString("buffered\n")
  checkError(err)

  fmt.Printf("wrote %d bytes\n", n4)
  w.Flush() // ensure all buffered ops have been applied

}
