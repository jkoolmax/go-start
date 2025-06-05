package main

import "fmt"


type Bird struct {
   name string
   canFly bool
   age int
}

func (ss Bird) String() string {
    return ss.name
}

func (b Bird) quack() string {
   if b.canFly == true && b.name =="duck" {
	return "can fly and quack"
   }
    return "can not fly"
}

func (b *Bird) getAge() int {
    return b.age
}

func main() {
  
   duck := Bird{name: "duck", canFly: true, age:2}
   turkey := Bird{name: "turkey", canFly: false, age: 1}
   
   //fmt.Println(duck) using Stringer e.g __str__
   
   fmt.Printf("%v age: %d, %v\n", duck.name, duck.getAge(), duck.quack()) 
   fmt.Printf("%v age: %d, %v\n", turkey.name, turkey.getAge(), turkey.quack())
 
}
