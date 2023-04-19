package main

import (
	"fmt"
	"reflect"
)

type General interface{}

// GData is holding General value
type GData interface {
	Set(name string, val General) GData
	Print()
}

type NData struct {
	Name string
	Data int
}

type GDataImpl struct {
	Name string
	Data General
}

func (nd *NData) Set(name string, val General) GData {
	nd.Name = name
	if reflect.TypeOf(val).Kind() == reflect.Int {
		nd.Data = val.(int)
	}
	return nd
}

func (nd *NData) Print() {
	fmt.Println(nd.Name, ":", nd.Data)
}

type SData struct {
	Name string
	Data string
}

func (sd *SData) Set(name string, val General) GData {
	sd.Name = name
	if reflect.TypeOf(val).Kind() == reflect.String {
		sd.Data = val.(string)
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Println(sd.Name, ":", sd.Data)
}

func main() {
	data := []GData{}
	data = append(data, new(NData).Set("num1", 1))
	data = append(data, new(SData).Set("num2", "aaaa"))
	data = append(data, new(NData).Set("num3", 3))
	data = append(data, new(SData).Set("num4", []string{"a", "b"}))

	for _, v := range data {
		v.Print()
	}
}
