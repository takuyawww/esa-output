package csv

import (
	"os"
	"reflect"
	"strconv"

	"github.com/takuyawww/esa-output/src/external"
)

const (
	outPutFileName = "esa.csv"

	separator = ","
	newLine   = "\n"
)

type Csv struct {
	Members   []*external.ResponseMembers
	Posts     []*external.ResponsePosts
	outputStr string
}

func New(
	members []*external.ResponseMembers,
	posts []*external.ResponsePosts,
) *Csv {
	return &Csv{
		Members: members,
		Posts:   posts,
	}
}

func (ec *Csv) String() *Csv {
	ec.outputStr = ec.HeaderString() + newLine + ec.bodyString()
	return ec
}

func (ec *Csv) Output() {
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

	_, err = f.Write([]byte(ec.outputStr))
	if err != nil {
		panic(err)
	}
}

func (ec *Csv) HeaderString() string {
	p := external.Post{}
	rp := reflect.ValueOf(p)
	rpt := rp.Type()

	var str string

	for i := 0; i < rpt.NumField(); i++ {
		str += reflect.TypeOf(p).Field(i).Tag.Get("headerLabel")

		if num := reflect.TypeOf(p).Field(i).Tag.Get("headerMultipleNum"); num != "" {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			for j := 1; j < n+1; j++ {
				str += reflect.TypeOf(p).Field(i).Tag.Get("headerLabel") + strconv.Itoa(j) + separator
			}
			// cutting last separator
			str = str[:len(str)-1]
		}
		str += separator
	}
	// cutting last separator
	return str[:len(str)-1]
}

func (ec *Csv) bodyString() string {
	var bodyStr string
	return bodyStr
}
