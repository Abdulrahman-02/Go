package operator

import "fmt"

func Operator(line1 chan string, line2 chan string, quit chan struct{}) {
	// Use an infinite loop to keep listening for new messages after handling a previous one
	for {
		// select blocks until one of the cases can be executed
		select {
		case msg := <-line1:
			// listen for incoming messages on line 1 and assign to msg variable
			fmt.Printf("Line1:%s\n", msg)
		case msg := <-line2:
			// listen for incoming messages on line2 and assign to msg variable
			fmt.Printf("Line 2: %s\n", msg)
		case <-quit:
			// listen for the quit channel to be closed and exit the function
			fmt.Println("Quit")
			return
		}
	}
}
