package main

import "fmt"

func main(){
	fmt.Println("This is my employee class which i have created in webstrom")
	s:= Employee{name:"Amit Sharma",age:20}
	fmt.Println(s.name)
	fmt.Println( s.age )
}

type  Employee struct{

	name string
	age int
}