package external

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	endpointFmt = EsaAPIOrigin +
		"/" +
		EsaAPIVersion +
		"/teams/%s/posts?access_token=%s&sort=%s&order=%s&per_page=%d&page=%d"
)

type PostsFetcher struct {
	qp       *QueryParams
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

func NewPostsFetcher(qp *QueryParams) *PostsFetcher {
	newTrue := true
	return &PostsFetcher{qp: qp, loopFlag: &newTrue}
}

func (f *PostsFetcher) Do() []*ResponsePosts {
	rps := make([]*ResponsePosts, 0)

	for *f.loopFlag {
		rp, err := do(f.qp)
		if err != nil {
			panic(err)
		}

		rps = append(rps, rp)
		if rp.NextPage == 0 {
			*f.loopFlag = false
		}
		f.qp.Page = rp.NextPage
	}

	return rps
}

func do(qp *QueryParams) (*ResponsePosts, error) {
	ep := buildEndpoint(endpointFmt, qp)

	res, err := http.Get(ep)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rp := new(ResponsePosts)

	if err := json.Unmarshal(b, rp); err != nil {
		return nil, err
	}

	return rp, nil
}
