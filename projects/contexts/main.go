package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("Doing something: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "masterkey")
	doAnother(anotherCtx)

	fmt.Printf("Doing something: myKey's value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("Doing another: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	// use context.TODO as empty (or starting) context placeholder when you're not sure which to use.
	//ctx := context.TODO()
	// use context.Background as empty (or starting) context placeholder when you intend to start a known context.
	ctx := context.Background()
	ctx = context.WithValue(ctx, "myKey", "hosis11223344")
	doSomething(ctx)
}
