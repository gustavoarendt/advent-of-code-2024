package main

import (
	"fmt"
	"os"

	"github.com/gustavoarendt/advent-of-code-2024/cmd"
)

func main() {
	if err := cmd.Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
