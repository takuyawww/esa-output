package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PostsFetcher struct {
	qp *QueryParams
}

type Post struct {
	Number         int           `json:"number"`
	Name           string        `json:"name"`
	FullName       string        `json:"full_name"`
	Wip            bool          `json:"wip"`
	BodyMd         string        `json:"body_md"`
	BodyHtml       string        `json:"body_html"`
	CreatedAt      time.Time     `json:"created_at"`
	Message        string        `json:"message"`
	Url            string        `json:"url"`
	UpdatedAt      time.Time     `json:"updated_at"`
	Tags           []string      `json:"tags"`
	Category       string        `json:"category"`
	RevisionNumber int           `json:"revision_number"`
	CreatedBy      PostCreatedBy `json:"created_by"`
	UpdatedBy      PostUpdatedBy `json:"updated_by"`
}

type PostCreatedBy struct {
	Myself     bool   `json:"myself"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type PostUpdatedBy struct {
	Myself     bool   `json:"myself"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Icon       string `json:"icon"`
}

type responsePosts struct {
	Posts      []Post `json:"posts"`
	PrevPage   int    `json:"prev_page"`
	NextPage   int    `json:"next_page"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	MaxPerPage int    `json:"max_per_page"`
}

const (
	endpointFmt = EsaAPIOrigin + "/" + EsaAPIVersion + "/teams/%s/posts?access_token=%s&sort=%s&order=%s&per_page=%d"
)

func NewPostsFetcher(qp *QueryParams) *PostsFetcher {
	return &PostsFetcher{qp: qp}
}

func (f *PostsFetcher) Do() {
	ep := buildEndpoint(endpointFmt, f.qp)
	res, err := http.Get(ep)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	rp := new(responsePosts)

	if err := json.Unmarshal(b, rp); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", rp)

	for _, post := range rp.Posts {
		fmt.Printf(
			"NAME: %-7s Category: %s CreatedBy: %s LastUpdatedBy: %s LastUpdatedAt: %s\n",
			post.Name,
			post.Category,
			post.CreatedBy.Name,
			post.UpdatedBy.Name,
			post.UpdatedAt,
		)
	}
}
