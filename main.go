package main

import (
	"os"

	"github.com/3at-developtment/sessionbeat/cmd"

	_ "github.com/3at-developtment/sessionbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
