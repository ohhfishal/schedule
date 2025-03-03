package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"

	"github.com/ohhfishal/schedule/cmd"
)

//go:embed sql/schema.sql
var schema string

func main() {
	if err := cmd.Run(context.Background(), os.Stdout, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
