package main

import (
	"io"
	"log"
	"os"

	"github.com/mokiat/preftime/prefix"
)

func main() {
	in := os.Stdin
	out := prefix.NewWriter(os.Stdout, prefix.TimestampFunc())

	if _, err := io.Copy(out, in); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
