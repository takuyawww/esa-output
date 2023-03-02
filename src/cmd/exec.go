package cmd

import (
	"errors"
	"fmt"

	// "runtime/debug"

	"github.com/alecthomas/kingpin/v2"
	"github.com/takuyawww/esa-output/src/external"
	"github.com/takuyawww/esa-output/src/output/csv"
)

func Exec() {
	printInformation()

	defer func() {
		if x := recover(); x != nil {
			// debug.PrintStack()
			errStr := fmt.Sprint(x)
			printErr(errors.New(errStr))
		}
	}()

	qp := parseFlag()

	esaPosts := external.NewPostsAPIFetcher(qp).Do()
	esaMembers := external.NewMembersAPIFetcher(qp).Do()

	csv := csv.New(esaMembers, esaPosts)
	csvStr := csv.String()
	csv.Output(csvStr)
}

func printInformation() {
	fmt.Println("********** Information **********")
	fmt.Println("Currently, esa accept up to 75 requests per user per 15 minutes.")
	fmt.Println("Operation is not guaranteed when the request limit is exceeded.")
	fmt.Println("See https://docs.esa.io/posts/102")
	fmt.Printf("**************************\n\n")
}

func printErr(err error) {
	fmt.Printf("error occurred, reason: %s", err.Error())
}

func parseFlag() *external.APIQueryParams {
	var (
		t  = kingpin.Flag("team", "your team name (*required*)").Required().String()
		at = kingpin.Flag("access-token", "your esa personal access token (*required*)").Required().String()
		pp = kingpin.Flag("per-page", "per page default 100 (in 20 ~ 100)").Default("100").Int()
		sp = kingpin.Flag("sort-posts", "sort by posts default number (sort by updated, created, number, stars, watches, comments, best_match)").Default("number").String()
		sm = kingpin.Flag("sort-members", "sort by members default joined (sort by posts_count, joined, last_accessed)").Default("number").String()
		o  = kingpin.Flag("order", "order by default asc (order by desc or asc)").Default("asc").String()
	)

	kingpin.Parse()

	return &external.APIQueryParams{
		Team:        *t,
		AccessToken: *at,
		SortPosts:   *sp,
		SortMembers: *sm,
		Order:       *o,
		PerPage:     *pp,
		Page:        1,
	}
}
