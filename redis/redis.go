package main

import (
	"bufio"
	"fmt"
	"os"

	// Import the Radix.v2 redis package.
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

func CreateUsersMap() map[int]string {
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

func getMapCollection() {
	names := CreateUsersMap()
	for k, v := range names {
		fmt.Printf("Key: %d Value: %s\n", k, v)
	}
}
func insert() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	// Importantly, use defer to ensure the connection is always properly
	// closed before exiting the main() function.
	defer conn.Close()

	// Send our command across the connection. The first parameter to Cmd()
	// is always the name of the Redis command (in this example HMSET),
	// optionally followed by any necessary arguments (in this example the
	// key, followed by the various hash fields and values).
	names := CreateUsersMap()
	for k, v := range names {
		resp := conn.Cmd("SET", k, v)
		fmt.Printf("Adding Value: %s\n", k, v)
		if resp.Err != nil {
			log.Fatal(resp.Err)
		}
	}
	//resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
	// Check the Err field of the *Resp object for any errors.

}
func del() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	resp := conn.Cmd("DEL", "2")
	if resp.Err != nil {
		log.Fatal(resp.Err)
	}
	fmt.Printf("Key:2 was deleted")
}
func get() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	resp := conn.Cmd("GET", "1")
	if resp.Err != nil {
		log.Fatal(resp.Err)
	}
	fmt.Print(resp)
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
	case 'i':
		insert()
		break
	case 'm':
		getMapCollection()
		break
	case 'd':
		del()
		break
	case 'g':
		get()
		break
	}
}
