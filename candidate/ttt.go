package main

import (
	"fmt"
	"time"
)

var ch chan bool

func aaa() {
	time.Sleep(1 * time.Second)
	ch <- true
}
func bbb() {

	id := <-ch
	fmt.Println(id)
}
func main() {
	ch = make(chan bool)
	fmt.Println(111)
	go bbb()
	go aaa()

	time.Sleep(5 * time.Second)
}
