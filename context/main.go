package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func exampleTimeout(parent context.Context, finder func() int, find func(context.Context, func() int)) {
	// Context that expires after 1 second. The cancel function is used to
	// release all resources associated with the context.
	ctx, cancel := context.WithTimeout(parent, time.Second)
	defer cancel()

	find(ctx, finder)
}

func exampleCancel(parent context.Context, finder func() int, find func(context.Context, func() int)) {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(parent)
	go func() {
		// time.Sleep(1 * time.Second)
		// cancelCancel()
		<-sig
		fmt.Println("received interrupt")
		cancel()
		time.Sleep(1)
		os.Exit(1)
	}()
	find(ctx, finder)
}

func exampleSignal(parent context.Context, finder func() int, find func(context.Context, func() int)) {

	// Different implementation for exampleCancel
	ctx, cancel := signal.NotifyContext(parent, os.Interrupt, syscall.SIGTERM)
	defer cancel()
	find(ctx, finder)
}

func main() {

	// Create a root context that is never cancelled and used in the main thread
	// to create children contexts.
	rootCtx := context.Background()

	finder := findSlow
	example := func() {
		exampleSignal(rootCtx, finder, find)
	}
	example()
}

func find(ctx context.Context, finder func() int) {

	// Stores the result of the find function
	res := make(chan int)
	go func() {
		res <- finder()
	}()

	// Wait for either of the following events to happen
	select {

	// Function finished before context expires
	case value := <-res:
		fmt.Println(value)
		return
	// Context finished before function
	case <-ctx.Done():
		fmt.Println("cleanup")
		fmt.Println(ctx.Err())
	}
}

func findFast() int {
	time.Sleep(100 * time.Millisecond)
	rand.Seed(time.Now().Unix())
	return 1
}
func findSlow() int {
	time.Sleep(2 * time.Second)
	rand.Seed(time.Now().Unix())
	return 2
}
