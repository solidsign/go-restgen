package fiber

import (
	"github.com/solidsign/go-restgen/codegen"
	"github.com/solidsign/go-restgen/commands/args"
	"github.com/solidsign/go-restgen/templates/swagger"
	"github.com/solidsign/go-restgen/utils"
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

func generateProtocolStructs(args args.Args) error {
	return codegen.New("protocol", "api/"+args.Group+"/protocol/"+args.EndpointName).
		Struct(utils.CapitalizeFirstLetter(args.Group) + utils.CapitalizeFirstLetter(args.EndpointName) + "Request").
		Struct(utils.CapitalizeFirstLetter(args.Group) + utils.CapitalizeFirstLetter(args.EndpointName) + "Response").
		Write()
}

func generateRouteInit(args args.Args) error {
	annotationString := ""
	if args.GenerateSwagger {
		annotationString = swagger.CreateAnnotationComments(swagger.RouteAnnotation{
			Summary:      args.Group + " " + args.EndpointName,
			Description:  "",
			Tags:         []string{args.Group},
			Produce:      "json",
			Secured:      args.Secured,
			Method:       args.Method,
			Url:          args.Url,
			EndpointName: args.EndpointName,
			Group:        args.Group,
			HasInput:     true,
		})
	}

	return codegen.New("routes", "api/"+args.Group+"/routes/"+args.EndpointName).
		Append(annotationString).
		FuncVoidStart("Init" + utils.CapitalizeFirstLetter(args.Group) + utils.CapitalizeFirstLetter(args.EndpointName)).
		FuncEnd().
		Write()
}

func generateController(args args.Args) error {
	return codegen.New("controller", "api/"+args.Group+"/controller/"+args.EndpointName).
		Import("github.com/gofiber/fiber/v2", args.ModuleName+"/api/"+args.Group+"/protocol").
		Struct(args.EndpointName+"Controller").
		FuncStart("Handle", "error", "ctx *fiber.Ctx").
		AppendLine("panic(\"Not implemented\")").
		FuncEnd().
		Write()
}

func generateUseCase(args args.Args) error {
	err := codegen.New("usecase", "api/"+args.Group+"/usecase/"+args.EndpointName).
		Interface(utils.CapitalizeFirstLetter(args.EndpointName) + "UseCase").
		Struct(args.EndpointName + "UseCaseImpl").
		Write()

	if err != nil {
		return err
	}

	return codegen.New("usecase", "api/"+args.Group+"/usecase/test_"+args.EndpointName).
		Import("github.com/stretchr/testify/assert", "testing").
		Struct(args.EndpointName+"UseCaseTest").
		FuncVoidStart("Test"+utils.CapitalizeFirstLetter(args.EndpointName), "t *testing.T").
		AppendLine("panic(\"not implemented\")").
		FuncEnd().
		Write()
}
