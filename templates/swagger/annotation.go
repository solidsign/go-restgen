package swagger

import (
	"fmt"
	"github.com/solidsign/go-restgen/utils"
	"strings"
)

type RouteAnnotation struct {
	Summary      string
	Description  string
	Tags         []string
	Produce      string
	Secured      bool
	Method       string
	Url          string
	EndpointName string
	Group        string
	HasInput     bool
}

// example:
// @Summary		Unsubscribe from user
// @Description	Stop receiving notifications of user's actions
// @Tags			subscriptions
// @Produce		json
// @Security		ApiKeyAuth
// @Param			input	body		protocol.SubscriptionRequest	true "input"
// @Success		200		{object}	protocol.SubscriptionResponse
// @Failure		400,401	{object}	protocol.ErrorResponse
// @Failure		500		{object}	protocol.ErrorResponse
// @Router			/api/unsubscribe [post]

func CreateAnnotationComments(annotation RouteAnnotation) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("// @Summary %s\n", annotation.Summary))
	b.WriteString(fmt.Sprintf("// @Description %s\n", annotation.Description))
	b.WriteString(fmt.Sprintf("// @Tags %s\n", strings.Join(annotation.Tags, ", ")))
	b.WriteString(fmt.Sprintf("// @Produce %s\n", annotation.Produce))
	if annotation.Secured {
		b.WriteString("// @Security ApiKeyAuth\n")
	}
	if annotation.HasInput {
		b.WriteString(fmt.Sprintf("// @Param input body protocol.%sRequest true \"input\"\n", utils.CapitalizeFirstLetter(annotation.Group)+utils.CapitalizeFirstLetter(annotation.EndpointName)))
	}
	b.WriteString(fmt.Sprintf("// @Success 200 {object} protocol.%sResponse\n", utils.CapitalizeFirstLetter(annotation.Group)+utils.CapitalizeFirstLetter(annotation.EndpointName)))
	b.WriteString("// @Failure 400,401 {object} protocol.ErrorResponse\n")
	b.WriteString("// @Failure 500 {object} protocol.ErrorResponse\n")
	b.WriteString(fmt.Sprintf("// @Router %s [%s]\n", annotation.Url, annotation.Method))

	return b.String()
}
