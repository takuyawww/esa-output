package external

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	postEndpointFmt = EsaAPIOrigin +
		"/" +
		EsaAPIVersion +
		"/teams/%s/posts?access_token=%s&sort=%s&order=%s&per_page=%d&page=%d"
)

type PostsFetcher struct {
	qp       *APIQueryParams
	loopFlag *bool
}

// comment out unused responses to reduce memory usage
type Post struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
	// FullName       string        `json:"full_name"`
	Wip bool `json:"wip"`
	// BodyMd         string        `json:"body_md"`
	// BodyHtml       string        `json:"body_html"`
	CreatedAt time.Time `json:"created_at"`
	// Message        string        `json:"message"`
	// Url            string        `json:"url"`
	UpdatedAt time.Time `json:"updated_at"`
	// Tags           []string      `json:"tags"`
	Category string `json:"category"`
	// RevisionNumber int           `json:"revision_number"`
	CreatedBy PostCreatedBy `json:"created_by"`
	UpdatedBy PostUpdatedBy `json:"updated_by"`
}

type PostCreatedBy struct {
	// Myself     bool   `json:"myself"`
	Name string `json:"name"`
	// ScreenName string `json:"screen_name"`
	// Icon       string `json:"icon"`
}

type PostUpdatedBy struct {
	// Myself     bool   `json:"myself"`
	Name string `json:"name"`
	// ScreenName string `json:"screen_name"`
	// Icon       string `json:"icon"`
}

type ResponsePosts struct {
	Posts []Post `json:"posts"`
	// PrevPage   int    `json:"prev_page"`
	NextPage int `json:"next_page"`
	// TotalCount int    `json:"total_count"`
	Page int `json:"page"`
	// PerPage    int    `json:"per_page"`
	// MaxPerPage int    `json:"max_per_page"`
}

func NewPostsFetcher(qp *APIQueryParams) *PostsFetcher {
	newTrue := true
	return &PostsFetcher{qp: qp, loopFlag: &newTrue}
}

func (f *PostsFetcher) Do() []*ResponsePosts {
	results := make([]*ResponsePosts, 0)

	for *f.loopFlag {
		result, err := f.do()
		if err != nil {
			panic(err)
		}

		results = append(results, result)
		if result.NextPage == 0 {
			*f.loopFlag = false
		}
		f.qp.Page = result.NextPage
	}

	return results
}

func (f *PostsFetcher) do() (*ResponsePosts, error) {
	ep := buildEndpoint(postEndpointFmt, f.qp, f.qp.SortPosts)

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
