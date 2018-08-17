package main

import "fmt"

type Address struct {
	Street, City, State, Zip string
	IsShippingAddress        bool
}
type Customer struct {
	FirstName, LastName, Email, Phone string
	Addresses                         []Address
}

func (c *Customer) ChangeEmail(newEmail string) {
	c.Email = newEmail
}
func (c Customer) ShippingAddress() string {
	for _, v := range c.Addresses {
		if v.IsShippingAddress == true {
			return fmt.Sprintf("%s, %s, %s, Zip - %s", v.Street, v.City, v.State,
				v.Zip)
		}
	}
	return ""
}
func (c *Customer) ToString() string {
	return fmt.Sprintf("Customer: %s %s, Email:%s", c.FirstName, c.LastName, c.Email)
}

func main() {
	c := Customer{
		FirstName: "Krysta",
		LastName:  "Boyle",
		Email:     "krysta@email.com",
		Phone:     "000-00-2923",
		Addresses: []Address{
			Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
			Address{
				Street: "49 Stevenson Street",
				City:   "San Francisco",
				State:  "CA",
				Zip:    "94105",
			},
		},
	}
	fmt.Println(c.ToString())
	c.ChangeEmail("krystab@gmail.com")
	fmt.Println("Customer after changing the Email:")
	fmt.Println(c.ToString())
	fmt.Println(c.ShippingAddress())
}
