package args

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Args struct {
	Method          string
	Url             string
	Group           string
	EndpointName    string
	Secured         bool
	GenerateSwagger bool
	GenerateUseCase bool
	Architecture    string
	Framework       string
	ModuleName      string
}

func ParseCmdArgs() (Args, error) {
	args := os.Args
	res := Args{
		Secured:         true,
		GenerateSwagger: true,
		GenerateUseCase: true,
		Architecture:    "cleanarch",
		Framework:       "fiber",
	}
	for i, arg := range args {
		if i == 0 {
			continue // skip program name
		}

		switch arg {
		case "--method":
			res.Method = normalizeMethod(args[i+1])
		case "--url":
			res.Url = normalizeUrl(args[i+1])
		case "--group":
			res.Group = args[i+1]
		case "--endpoint-name":
			res.EndpointName = args[i+1]
		case "--unsecured":
			res.Secured = false
		case "--no-swagger":
			res.GenerateSwagger = false
		case "--no-usecase":
			res.GenerateUseCase = false
		case "--architecture":
			res.Architecture = args[i+1]
		case "--framework":
			res.Framework = args[i+1]
		case "--module-name":
			res.ModuleName = args[i+1]
		default:
			switch i {
			case 1:
				res.Method = normalizeMethod(arg)
			case 2:
				res.Url = normalizeUrl(arg)
			case 3:
				res.Group = arg
			case 4:
				res.EndpointName = arg
			}
		}
	}

	res = fillGroupAndEndpoint(res)
	res = fillModuleName(res)

	return res, validateArgs(res)
}

func fillGroupAndEndpoint(args Args) Args {
	if args.Group != "" && args.EndpointName != "" {
		return args
	}

	url := strings.Split(args.Url, "/")
	args.EndpointName = url[len(url)-2]
	if len(url) > 2 {
		args.Group = url[len(url)-3]
	} else {
		args.Group = "unnamed"
	}

	return args
}

func fillModuleName(args Args) Args {
	if args.ModuleName != "" {
		return args
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return args
	}

	file, err := os.OpenFile(dir+"/go.mod", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return args
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "module") {
			args.ModuleName = strings.TrimPrefix(text, "module ")
			return args
		}
	}
	return args
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

func normalizeMethod(method string) string {
	return strings.ToLower(method)
}

func normalizeUrl(url string) string {
	url = strings.ToLower(url)
	url = strings.ReplaceAll(url, "\\", "/")
	if strings.HasSuffix(url, "/") {
		return url
	}
	return url + "/"
}
