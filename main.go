package main

import (
	"fmt"
	"sample/hello"
	"strconv"
)

type MyData struct {
	Name string
	Data []int
}

type intp int

func (num intp) IsPrime() bool {
	for i := 2; i < int(num); i++ {
		if int(num)%i == 0 {
			return false
		}
	}
	return true
}

// function which extract prime factor from number
func (num intp) PrimeFactor() []int {
	result := []int{}
	x := int(num)

	for i := 2; i < int(num); i++ {
		if x%i == 0 {
			x /= i
			result = append(result, i)
		}
	}
	return result
}

func (num *intp) maxPrime() {
	pf := num.PrimeFactor()
	*num = intp(pf[len(pf)-1])
}

func main() {
	s := hello.Input("Type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n)
	fmt.Printf("%d is prime? %t.\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x.maxPrime()
	fmt.Printf("%d is prime? %t.\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
	x++
	fmt.Printf("%d is prime? %t.\n", x, x.IsPrime())
	fmt.Println(x.PrimeFactor())
}
