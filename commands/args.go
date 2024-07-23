package commands

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Args struct {
	Method          string
	Url             string
	Secured         bool
	GenerateSwagger bool
}

func ParseCmdArgs() (Args, error) {
	args := os.Args
	res := Args{
		Secured:         true,
		GenerateSwagger: true,
	}
	for i, arg := range args {
		if i == 0 {
			continue // skip program name
		}

		switch arg {
		case "--method":
			res.Method = strings.ToLower(args[i+1])
		case "--url":
			res.Url = args[i+1]
		case "--unsecured":
			res.Secured = false
		case "--no-swagger":
			res.GenerateSwagger = false
		default:
			switch i {
			case 1:
				res.Method = strings.ToLower(arg)
			case 2:
				res.Url = arg
			}
		}
	}

	return res, validateArgs(res)
}

func validateArgs(args Args) error {
	if args.Method == "" {
		return fmt.Errorf("missing method")
	}
	if args.Url == "" {
		return fmt.Errorf("missing url")
	}
	possibleMethods := []string{"get", "post", "put", "delete"}
	if slices.Contains(possibleMethods, args.Method) == false {
		return fmt.Errorf("invalid method. Possible values: %v", possibleMethods)
	}

	return nil
}
