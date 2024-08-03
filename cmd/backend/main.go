package main

import (
	"context"
	"fmt"
	"os"
	"vdo-be/internal/handler"
)

func main() {
	fmt.Println("starting server")
	ctx := context.Background()

	if err := handler.RunServer(ctx, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
