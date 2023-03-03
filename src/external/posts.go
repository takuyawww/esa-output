package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	postsAPIEndpointFmt = EsaAPIOrigin +
		"/" +
		EsaAPIVersion +
		"/teams/%s/posts?access_token=%s&sort=%s&order=%s&per_page=%d&page=%d"
)

type PostsAPIFetcher struct {
	qp       *APIQueryParams
	loopFlag *bool
}

type Post struct {
	Number              int           `json:"number" headerLabel:"ID"`
	Name                string        `json:"name" headerLabel:"Title"`
	Category            string        `json:"category" headerLabel:"Category" headerMultipleNum:"5"`
	CreatedBy           PostCreatedBy `json:"created_by" headerLabel:"CreatedBy"`
	IsActiveUserCreated bool          `json:"-" headerLabel:"IsActiveUserCreated"`
	CreatedAt           time.Time     `json:"created_at" headerLabel:"CreatedAt"`
	UpdatedBy           PostUpdatedBy `json:"updated_by" headerLabel:"LastUpdatedBy"`
	UpdatedAt           time.Time     `json:"updated_at" headerLabel:"LastUpdatedAt"`
	Wip                 bool          `json:"wip" headerLabel:"WIP"`
	// FullName       string        `json:"full_name"`
	// BodyMd         string        `json:"body_md"`
	// BodyHtml       string        `json:"body_html"`
	// Message        string        `json:"message"`
	// Url            string        `json:"url"`
	// Tags           []string      `json:"tags"`
	// RevisionNumber int           `json:"revision_number"`
}

type PostCreatedBy struct {
	Name string `json:"name"`
	// Myself     bool   `json:"myself"`
	// ScreenName string `json:"screen_name"`
	// Icon       string `json:"icon"`
}

type PostUpdatedBy struct {
	Name string `json:"name"`
	// Myself     bool   `json:"myself"`
	// ScreenName string `json:"screen_name"`
	// Icon       string `json:"icon"`
}

type ResponsePosts struct {
	Posts    []Post `json:"posts"`
	NextPage int    `json:"next_page"`
	Page     int    `json:"page"`
	// PrevPage   int    `json:"prev_page"`
	// TotalCount int    `json:"total_count"`
	// PerPage    int    `json:"per_page"`
	// MaxPerPage int    `json:"max_per_page"`
}

func NewPostsAPIFetcher(qp *APIQueryParams) *PostsAPIFetcher {
	newTrue := true
	return &PostsAPIFetcher{qp: qp, loopFlag: &newTrue}
}

func (f *PostsAPIFetcher) Do() []Post {
	results := make([]Post, 0)

	for *f.loopFlag {
		result, err := f.do()
		if err != nil {
			panic(err)
		}

		results = append(results, result.Posts...)

		if result.NextPage == 0 {
			*f.loopFlag = false
		}
		f.qp.Page = result.NextPage
	}

	return results
}

func (f *PostsAPIFetcher) do() (*ResponsePosts, error) {
	ep := buildAPIEndpoint(postsAPIEndpointFmt, f.qp, f.qp.SortPosts)

	res, err := http.Get(ep)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := new(ResponsePosts)

	if err := json.Unmarshal(b, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (p Post) ReflectValueToString(fieldName string, members []Member) string {
	switch fieldName {
	case "Number":
		return strconv.Itoa(p.Number)
	case "Name":
		return p.Name
	case "CreatedBy":
		return p.CreatedBy.Name
	case "CreatedAt":
		return p.CreatedAt.Format("2006/01/02")
	case "IsActiveUserCreated":
		var isActiveUserCreated bool

		for _, m := range members {
			if m.Name == p.CreatedBy.Name {
				isActiveUserCreated = true
				break
			}
		}

		return fmt.Sprintf("%t", isActiveUserCreated)
	case "UpdatedBy":
		return p.UpdatedBy.Name
	case "UpdatedAt":
		return p.UpdatedAt.Format("2006/01/02")
	case "Wip":
		return fmt.Sprintf("%t", p.Wip)
	default:
		return ""
	}
}

func (p Post) ReflectValueMultipleToString(fieldName string, index int) string {
	switch fieldName {
	case "Category":
		split := strings.Split(p.Category, "/")
		if len(split) > index {
			return split[index]
		}
		return ""
	default:
		return ""
	}
}
