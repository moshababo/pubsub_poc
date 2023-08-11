package main

import (
	"fmt"
	"os"
)

func main() {
	if err := runServer(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
