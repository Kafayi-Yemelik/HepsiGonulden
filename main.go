package main

import (
	"HepsiGonulden/cmd"
	"HepsiGonulden/pkg"
	"fmt"
)

func main() {

	err := pkg.Init()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	cmd.NewCommand().Execute()

}
