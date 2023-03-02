package csv

import (
	"os"
	"reflect"
	"strconv"

	"github.com/takuyawww/esa-output/src/external"
)

const (
	outPutFileName = "esa.csv"

	separator = "\t"
	newLine   = "\n"
)

type Csv struct {
	Members []external.Member
	Posts   []external.Post
}

func New(
	members []external.Member,
	posts []external.Post,
) *Csv {
	return &Csv{
		Members: members,
		Posts:   posts,
	}
}

func (ec *Csv) String() string {
	return ec.HeaderString() + newLine + ec.bodyString()
}

func (ec *Csv) Output(str string) {
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

	_, err = f.Write([]byte(str))
	if err != nil {
		panic(err)
	}
}

func (ec *Csv) HeaderString() string {
	p := external.Post{}
	pv := reflect.ValueOf(p)
	pt := pv.Type()

	var str string

	for i := 0; i < pt.NumField(); i++ {
		num := reflect.TypeOf(p).Field(i).Tag.Get("headerMultipleNum")

		if num == "" {
			str += reflect.TypeOf(p).Field(i).Tag.Get("headerLabel")
		} else {
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
	var str string

	for _, post := range ec.Posts {
		pv := reflect.ValueOf(post)
		pt := pv.Type()

		for i := 0; i < pt.NumField(); i++ {
			field := pt.Field(i)
			value := post.ReflectValueToString(field.Name)
			str += value + separator
		}

		str += newLine
	}

	return str[:len(str)-1]
}

// TODO
// カテゴリーを/区切りで埋めていく
// 削除済みユーザーかどうかを捜査する
