package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			panic(err)
		}
		line := scanner.Text()
		timestamp := time.Now()
		fmt.Printf("[")
		fmt.Printf(timestamp.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("] ")
		fmt.Println(line)
	}
}
