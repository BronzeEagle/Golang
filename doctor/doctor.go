package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Customer struct {
	Name   string
	RoleID int
}

func main() {
	fmt.Println("Starting the application...")
	url := "http://localhost:8090/getusingle?name=Aline"
	//http://localhost:8090/getuserbyrol
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var customer Customer
		err = json.Unmarshal(data, &customer)
		if err != nil {
			fmt.Println("There was an error:", err)
			return
		}
		fmt.Printf("Name:%s, Role:%d \n", customer.Name, customer.RoleID)
	}
	fmt.Println("Terminating the application...")
}
