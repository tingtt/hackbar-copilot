package main

import (
	"log/slog"
	"os"
)

func main() {
	err := run()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	return nil
}
