package csv

import (
	"os"

	"github.com/takuyawww/esa-csv/src/external"
)

const (
	outPutFileName = "esa-csv.csv"
)

type EsaCsv struct {
	Members []*external.ResponseMembers
	Posts   []*external.ResponsePosts
	csvStr  string
}

func NewEsaCsv(
	members []*external.ResponseMembers,
	posts []*external.ResponsePosts,
) *EsaCsv {
	return &EsaCsv{
		Members: members,
		Posts:   posts,
	}
}

func (ec *EsaCsv) ToCsvString() *EsaCsv {
	ec.markIsActiveUserCreatedPost()

	ec.csvStr = ec.toHeaderCsvString() + "\n" + ec.toBodyCsvString()

	return ec
}

func (ec *EsaCsv) OutputCsv() {
	f, err := os.Create(outPutFileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err = f.Write([]byte(ec.csvStr))
	if err != nil {
		panic(err)
	}
}

func (ec *EsaCsv) markIsActiveUserCreatedPost() {

}

func (ec *EsaCsv) toHeaderCsvString() string {
	var headerStr string
	return headerStr
}

func (ec *EsaCsv) toBodyCsvString() string {
	var bodyStr string
	return bodyStr
}
