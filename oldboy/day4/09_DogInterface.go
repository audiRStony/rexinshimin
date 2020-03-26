package main

/*
@Time : 2020-03-25 14:07
@Author : audiRStony
@File : 09_DogInterface.go
@Software: GoLand
*/

import (
    "fmt"
)

type Animal interface {
    Bark() string
    Walk() string
}


type Dog struct {
    name string
}

func (dog Dog) Bark() string  {
    fmt.Println(dog.name + ":wang wang wang")
    return "wang wang wang"
}

func (dog Dog) Walk() string  {
    fmt.Println(dog.name + ": walk to park")
    return "walk to park"
}

func main() {
    var animal   Animal

    fmt.Println("animal value is :",animal)
    fmt.Printf("animal type is : %T\n",animal)

    animal = Dog{"旺财"}
    animal.Bark()
    animal.Walk()

    fmt.Println("animal value is:",animal)
    fmt.Printf("animal type is %T\n",animal)

}

