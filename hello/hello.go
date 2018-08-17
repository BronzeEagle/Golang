package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func getName(name string) string {
	return "Hello " + name + "Boyle "
}

func main() {
	fmt.Printf(getName("Krysta ") + " \n")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8091", nil))
}
