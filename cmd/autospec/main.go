package main

import (
	"fmt"
	"os"

	"github.com/Jeielsantosdev/autospec/internal/autospec"
)

func main() {
	app := autospec.New()

	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}