package utility

import (
	"fmt"
)

func PrintUserName(name string) string {
	return "Hello " + name
}
func Testing() {
	fmt.Println("Krysta the testing is working ")
}
func CreateMap() map[int]string {
	// Declare and initialize map using map literal
	names := map[int]string{
		1: "Aline",
		2: "Krysta",
		3: "Samantha",
		4: "Lisa",
		5: "Sara",
		6: "Miranda",
	}
	return names
}
