package main

import (
	"context"
	"fmt"
	"time"
)

func listener(ctx context.Context, i int) {
	fmt.Printf("listner %d is waiting\n", i)

	// this will block until the context given is cancelled
	<-ctx.Done()

	fmt.Printf("listener %d is exiting \n", i)
}

func main() {
	// create a background context
	ctx := context.Background()

	// to cancel it
	ctx, cancel := context.WithCancel(ctx)

	// defer cancellation of the context
	// to ensure that any ressources are cleaned up regardless of how the function exits
	defer cancel()

	// create 5 listeners
	for i := 0; i < 5; i++ {
		go listener(ctx, i)
	}

	// allow the listeners to start
	time.Sleep(time.Millisecond * 500)

	fmt.Println("canceling the context")

	// cancel the context adn tell the listeners to exit
	cancel()

	// allow the listeners to exit
	time.Sleep(time.Millisecond * 500)
}
