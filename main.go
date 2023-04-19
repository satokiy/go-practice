package main

import (
	"fmt"
	"strconv"
	"time"
)

func hello(s string, t int, n int) {
	for i := 1; i < n; i++ {
		fmt.Printf("<%d %s>", i, s)
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
}

func prmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n int, c chan string) {
	const name string = "first-"
	for i := 0; i < n; i++ {
		s := name + strconv.Itoa(i)
		prmsg(n, s)
		c <- s
	}
}

func second(n int, c chan string) {
	for i := 0; i < 10; i++ {
		prmsg(n, "second-["+<-c+"]")
	}
}

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}

func main() {
	c := make(chan string)
	go first(10, c)
	second(10, c)
	fmt.Println("end")
}
