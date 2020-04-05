package main

/*
@Time : 2020-03-25 11:21
@Author : audiRStony
@File : 08_exampleInterface.go
@Software: GoLand
*/
import (
    "fmt"
)

type Shaper interface {
    Area() float32
}
type Square struct {
    side float32
}

func (sq *Square) Area() float32  {
    return sq.side * sq.side
}

func main()  {


    sq1 := new(Square)
    sq1.side = 5
    // areaIntf := sq1
    // fmt.Printf("面积为 %f\n",areaIntf.Area())
    areaIntf := Shaper(sq1)
    fmt.Printf("面积为 %f\n",areaIntf.Area())
}

