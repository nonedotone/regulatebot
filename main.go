package main

import (
	"context"
)

func main() {
	printVersion()
	checkFlag()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	h := NewHandler(ctx, token)
	h.botUpdate()
}
