package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ohhfishal/schedule/cmd"
)

func main() {
	if err := cmd.Run(context.Background(), os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
