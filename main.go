package main

import (
	"fmt"
)

type Data interface{
	Initial(name string, data []int)
	PrintData()
} 

type MyData struct {
	Name string
	Data []int
}

func (md *MyData) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

func (md *MyData) Check() {
	fmt.Printf("Check!! [%s]", md.Name)
}

// PrintData
func (md *MyData) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}


func main() {
	var ob Data = new(MyData)
	ob.Initial("Sachiko", []int{1,2,3})
	ob.PrintData()
}
