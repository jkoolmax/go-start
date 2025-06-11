 package main

 import (
    "bufio"
    "fmt"
    "io"
    "os"
 )

 // helper to streamline error check
 func checkError(e error) {
   if e != nil {
     panic(e)
   }
 }

 func main() {

     // dumping the whole file into the memory
     //dat, err := os.ReadFile("/tmp/text.txt")

     //checkError(err)
     //fmt.Print(string(dat))

     // read a specific part on a file
     f, err := os.Open("/tmp/text.txt")

     checkError(err)
     
     // allow 5 bytes to be read
     b1 := make([]byte, 5)
     n1, err := f.Read(b1)

     checkError(err)
     fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) // read the first 5bytes

     o2, err := f.Seek(6, io.SeekStart)
     checkError(err)
     b2 := make([]byte, 2)
     n2, err := f.Read(b2)

     checkError(err)
     fmt.Printf("%d bytes @ %d: ", n2, o2)
     fmt.Printf("%v\n", string(b2[:n2]))
    
     // seeking relative the current cursor
     _, err = f.Seek(2, io.SeekCurrent)
     checkError(err)
    
     // seeking relative to the end of the file
     _, err = f.Seek(-4, io.SeekEnd)
     checkError(err)

     o3, err := f.Seek(6, io.SeekStart)
     checkError(err)
     b3 := make([]byte, 2)
     n3, err := io.ReadAtLeast(f, b3, 2)

     checkError(err)
     fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

     _, err = f.Seek(0, io.SeekStart)
     checkError(err)

     r4 := bufio.NewReader(f)
     b4, err := r4.Peek(5)

     checkError(err)
     fmt.Printf("5 bytes: %s\n", string(b4))

     f.Close()


 }
