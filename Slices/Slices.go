package main

import (
	"bufio"
	"fmt"
	"os"
)

func printPartialSlice() {
	a := [8]string{"A", "K", "r", "y", "s", "t", "a", "z"}
	var b []string = a[1:7] //creates a slice from a[1] to a[3]
	fmt.Println(b)
	fmt.Printf("length of slice %d capacity %d", len(b), cap(b))
}
func printAllSlice() {
	names := []string{"Aline", "Krysta", "Sara", "Miranda"}
	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
func createMap() map[int]string {
	// Declare and initialize map using map literal
	names := map[int]string{
		1: "Aline",
		2: "Krysta",
		3: "Samantha",
		4: "Lisa",
	}
	return names
}

//Is shared as repository
func PrintMap() {
	names := createMap()
	for k, v := range names {
		fmt.Printf("Key: %d Value: %s\n", k, v)
	}
}

func swap(x, y string) (string, string) {
	return y, x
}
func printSwap() {
	a, b := swap("Krysta", "Aline")
	fmt.Println(a, b)
}
func main() {

	fmt.Println("SIM Please input option")
	fmt.Println("---------------------")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case 'a':
		printAllSlice()
		break
	case 'b':
		createMap()
		break
	case 'c':
		printPartialSlice()
		break
	case 'm':
		PrintMap()
		break
	case 's':
		printSwap()
		break
	}
}
