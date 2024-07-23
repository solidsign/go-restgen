package commands

import "restgen/commands/args"

type ExecuteMethod func(args args.Args) error
