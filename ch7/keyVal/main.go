package main

import (
	"context"
	"fmt"
)

type aKey string

func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	// found value: secretValue
	// key not found: wrongVal
	
	key := aKey("secret-key")
	ctx := context.WithValue(context.Background(), key, "secretValue")

	searchKey(ctx, key)
	searchKey(ctx, aKey("wrongVal"))
}
