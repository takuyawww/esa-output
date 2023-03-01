package cmd

import (
	"fmt"

	"github.com/alecthomas/kingpin/v2"
	"github.com/takuyawww/esa-csv/src/external"
)

func Exec() {
	printInformation()

	qp := parseFlag()
	external.NewPostsFetcher(qp).Do()
}

func printInformation() {
	fmt.Println("********** Information **********")
	fmt.Println("Currently, esa accept up to 75 requests per user per 15 minutes.")
	fmt.Println("Operation is not guaranteed when the request limit is exceeded.")
	fmt.Println("See https://docs.esa.io/posts/102")
	fmt.Printf("**************************\n\n")
}

func parseFlag() *external.QueryParams {
	var (
		t  = kingpin.Flag("team", "your team name (*required*)").Required().String()
		at = kingpin.Flag("access-token", "your esa personal access token (*required*)").Required().String()
		pp = kingpin.Flag("per-page", "per page default 100 (in 20 ~ 100)").Default("100").Int()
		s  = kingpin.Flag("sort", "sort at default number (sort by updated, created, number, stars, watches, comments, best_match)").Default("number").String()
		o  = kingpin.Flag("order", "order by default asc (order by desc or asc)").Default("asc").String()
	)

	kingpin.Parse()

	return &external.QueryParams{
		Team:        *t,
		AccessToken: *at,
		PerPage:     *pp,
		Sort:        *s,
		Order:       *o,
	}
}
