package external

import (
	"fmt"
)

const (
	EsaAPIOrigin  = "https://api.esa.io"
	EsaAPIVersion = "v1"
)

type QueryParams struct {
	Team        string
	AccessToken string
	Sort        string
	Order       string
	PerPage     int
}

func buildEndpoint(
	endpointFmt string,
	qp *QueryParams,
) string {
	return fmt.Sprintf(endpointFmt, qp.Team, qp.AccessToken, qp.Sort, qp.Order, qp.PerPage)
}
