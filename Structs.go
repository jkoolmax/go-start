package main

import "fmt"

type person struct {
   name string
   age int
}

func newPerson(name string) *person {
   
    p := person{name: name}
    p.age = 39
    return &p

}

func main(){
  fmt.Println(person{"Bob", 20})
  fmt.Println(person{name: "Alice", age: 30})
  fmt.Println(person{name: "Fred"})

  fmt.Println(&person{name: "Ann", age:40})
  fmt.Println(newPerson("John"))

  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  sp.age = 52
  fmt.Println(sp.age)

  dog := struct {
	  name string
	  isGood bool
      }{
        "Rex", false,
      }
   
   fmt.Println(dog)
}
