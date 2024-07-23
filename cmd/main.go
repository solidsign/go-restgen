package main

import (
	"fmt"
	"restgen/commands"
	"restgen/commands/args"
)

func main() {
	args, err := args.ParseCmdArgs()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(args)

	execute, err := commands.GetExecutor(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = execute(args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
