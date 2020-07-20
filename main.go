package main

import (
	"fmt"
	"os"

	commands "github.com/rbreeze/zoom-cli/cmd"
)

func main() {
	if err := commands.NewCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
