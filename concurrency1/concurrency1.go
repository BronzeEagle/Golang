package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func getName(name string) string {
	return "Hello " + name + "Boyle "
}
func main() {
	go numbers()
	go alphabets()
	fmt.Printf(getName("Krysta ") + " \n")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
