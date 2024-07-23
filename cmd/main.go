package main

import (
	"fmt"
	"github.com/solidsign/go-restgen/commands"
	"github.com/solidsign/go-restgen/commands/args"
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
