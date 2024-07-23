package fiber

import (
	"github.com/solidsign/go-restgen/codegen"
	"github.com/solidsign/go-restgen/commands/args"
	"strings"
)

func Execute(args args.Args) error {
	if err := generateProtocolStructs(args); err != nil {
		return err
	}
	if err := generateRouteInit(args); err != nil {
		return err
	}
	if err := generateController(args); err != nil {
		return err
	}
	if args.GenerateUseCase {
		if err := generateUseCase(args); err != nil {
			return err
		}
	}
	return nil
}

func capitalizeFirstLetter(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

func generateProtocolStructs(args args.Args) error {
	return codegen.New("protocol", "api/"+args.Group+"/protocol/"+args.EndpointName).
		Struct(capitalizeFirstLetter(args.Group) + capitalizeFirstLetter(args.EndpointName) + "Request").
		Struct(capitalizeFirstLetter(args.Group) + capitalizeFirstLetter(args.EndpointName) + "Response").
		Write()
}

func generateRouteInit(args args.Args) error {
	return codegen.New("routes", "api/"+args.Group+"/routes/"+args.EndpointName).
		FuncVoidStart("Init" + capitalizeFirstLetter(args.Group) + capitalizeFirstLetter(args.EndpointName)).
		FuncEnd().
		Write()
}

func generateController(args args.Args) error {
	return codegen.New("controller", "api/"+args.Group+"/controller/"+args.EndpointName).
		Import("github.com/gofiber/fiber/v2", args.ProjectName+"/api/"+args.Group+"/protocol").
		Struct(args.EndpointName+"Controller").
		FuncStart("Handle", "error", "ctx *fiber.Ctx").
		AppendLine("panic(\"Not implemented\")").
		FuncEnd().
		Write()
}

func generateUseCase(args args.Args) error {
	err := codegen.New("usecase", "api/"+args.Group+"/usecase/"+args.EndpointName).
		Interface(capitalizeFirstLetter(args.EndpointName) + "UseCase").
		Struct(args.EndpointName + "UseCaseImpl").
		Write()

	if err != nil {
		return err
	}

	return codegen.New("usecase", "api/"+args.Group+"/usecase/test_"+args.EndpointName).
		Import("github.com/stretchr/testify/assert", "testing").
		Struct(args.EndpointName+"UseCaseTest").
		FuncVoidStart("Test"+capitalizeFirstLetter(args.EndpointName), "t *testing.T").
		AppendLine("panic(\"not implemented\")").
		FuncEnd().
		Write()
}
