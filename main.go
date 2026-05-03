package main

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/mokiat/preftime/prefix"
)

func main() {
	in := os.Stdin
	out := prefix.NewWriter(os.Stdout, prefix.TimestampFunc())

	_, err := io.Copy(out, in)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalf("Error: %v", err)
	}
}
