package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Println("context type:\t%T\n", ctx)
	fmt.Println("-------")

	ctx, cancel :=  context.WithCancel(ctx)
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Println("cancel:\t\t", cancel)
	fmt.Println("cancel type:\t%T\n", cancel)
	fmt.Println("-------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Println("cancel:\t\t", cancel)
	fmt.Println("cancel type:\t%T\n", cancel)
	fmt.Println("-------")

}
