package main

import (
	"bufio"
	"fmt"
	"os"
	"utility"
)

func printName() {
	utility.Testing()
}
func getMapCollection() {
	names := utility.CreateMap()
	for k, v := range names {
		fmt.Printf("Key: %d Value: %s\n", k, v)
	}
}

//*** Example 1
func sendEmail(isSent chan bool) {
	fmt.Println("Sending Email to Krysta by goroutine")
	isSent <- true
}
func channelExample1() {
	isSent := make(chan bool)
	go sendEmail(isSent)
	<-isSent //Blocks until get some data and then go to next line
	fmt.Println("Email was sent")
}

////*** END Example 1

func stringProducer() {
	messages := make(chan string)
	go func() { messages <- "Sending Msg to a channel" }()
	msg := <-messages
	fmt.Println(msg)
}

//**** Check Channel Is Closed
func producerString(chnlNames chan string) {
	names := utility.CreateMap()
	for _, v := range names {
		chnlNames <- v
	}
	close(chnlNames)
}
func checkChannelIsClosed() {
	ch := make(chan string)
	go producerString(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

func buffering() {
	messages := make(chan string, 3)
	messages <- "Aline"
	messages <- "Krysta"
	messages <- "Samantha"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

func lexicalConfinement() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}
	results := chanOwner()
	consumer(results)
}

//***************
func main() {
	fmt.Println("SIM Please input option")
	fmt.Println("---------------------")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(char)

	switch char {
	case 'a':
		channelExample1()
		break
	case 'c':
		stringProducer()
		break
	case 'd':
		buffering()
		break
	case 'e':
		fmt.Println("Check String Channel Has Been Closed")
		checkChannelIsClosed()
		break
	case 'm':
		getMapCollection()
		break

	case 'z':
		fmt.Println("Testing")
		printName()
		break
	case 'l':
		lexicalConfinement()
	}
}
