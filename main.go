package main

import (
	"HepsiGonulden/cmd"
	"HepsiGonulden/config"
	"fmt"
)

func main() {

	err := config.Init()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cmd.NewCommand().Execute()

}
