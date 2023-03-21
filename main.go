package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mokiat/preftime/prefix"
)

func main() {
	in := os.Stdin
	out := prefix.NewWriter(os.Stdout, prefix.TimestampFunc())

	_, err := io.Copy(out, in)
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "preftime error: %s\n", err)
		os.Exit(1)
	}
}
