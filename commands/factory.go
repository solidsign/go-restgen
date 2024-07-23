package commands

import (
	"fmt"
	"restgen/commands/args"
	cleanArchFiber "restgen/templates/cleanarch/fiber"
)

func GetExecutor(args args.Args) (ExecuteMethod, error) {
	switch args.Architecture {
	case "cleanarch":
		switch args.Framework {
		case "fiber":
			return cleanArchFiber.Execute, nil
		default:
			return nil, fmt.Errorf("unknown framework %s for arch %s", args.Framework, args.Architecture)
		}
	default:
		return nil, fmt.Errorf("unknown architecture: %s", args.Architecture)
	}
}
