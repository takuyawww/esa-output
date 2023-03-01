package external

import (
	"fmt"
)

const (
	EsaAPIOrigin  = "https://api.esa.io"
	EsaAPIVersion = "v1"
)

type APIQueryParams struct {
	Team        string
	AccessToken string
	SortPosts   string
	SortMembers string
	Order       string
	PerPage     int
	Page        int
}

func buildEndpoint(
	endpointFmt string,
	qp *APIQueryParams,
	sortBy string,
) string {
	return fmt.Sprintf(
		endpointFmt,
		qp.Team,
		qp.AccessToken,
		sortBy,
		qp.Order,
		qp.PerPage,
		qp.Page,
	)
}
