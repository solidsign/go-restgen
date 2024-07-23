package main

import (
	"fmt"
	"restgen/commands"
)

func main() {
	args, err := commands.ParseCmdArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(args)
}
