package main

import (
	"fmt"
	"os"

	"github.com/MohammadTaghipour/flumint/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		_, err = fmt.Fprintln(os.Stderr, "Error:", err)
		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}
}
