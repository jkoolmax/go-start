package main

import (
	"fmt"
	"reflect"
)

type Person struct {
  name string
  postname string
}

func intSeq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func handlePerson(persons ...Person) func() *Person {
        var prsn Person
	return func() *Person {
          for _, person := range persons {
              if person.name != "josue" {
                  continue 
		}
	       prsn = person
	  }
       fmt.Println(reflect.TypeOf(prsn).String() == "main.Person")
       return &prsn
   }
}

func main() {
  
   var persons []Person	
   peopleData := map[string]string{"josue": "kula",
   "kool": "max"}

  for key, name := range peopleData {
          p := Person{key, name}
	  persons = append(persons, p)
  }

  fmt.Println("persons: ", persons)
  
  prsns := handlePerson(persons...)
  fmt.Println(prsns().name)
}

