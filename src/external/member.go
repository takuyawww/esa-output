package external

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	memberEndpointFmt = EsaAPIOrigin +
		"/" +
		EsaAPIVersion +
		"/teams/%s/members?access_token=%s&sort=%s&order=%s&per_page=%d&page=%d"
)

type MembersFetcher struct {
	qp       *APIQueryParams
	loopFlag *bool
}

// comment out unused responses to reduce memory usage
type Member struct {
	// Myself 				bool   `json:"myself"`
	Name string `json:"name"`
	// ScreenName 		string `json:"screen_name"`
	// Icon 					string `json:"icon"`
	// Role 					string `json:"role"`
	// PostsCount 		int    `json:"posts_count"`
	// JoinedAt 			time.Time `json:"joined_at"`
	// LastAccessedAt time.Time `json:"last_accessed_at"`
}

type ResponseMembers struct {
	Members []Member `json:"members"`
	// PrevPage   int    `json:"prev_page"`
	NextPage int `json:"next_page"`
	// TotalCount int    `json:"total_count"`
	Page int `json:"page"`
	// PerPage    int    `json:"per_page"`
	// MaxPerPage int    `json:"max_per_page"`
}

func NewMembersFetcher(qp *APIQueryParams) *MembersFetcher {
	newTrue := true
	return &MembersFetcher{qp: qp, loopFlag: &newTrue}
}

func (f *MembersFetcher) Do() []*ResponseMembers {
	results := make([]*ResponseMembers, 0)

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

func (f *MembersFetcher) do() (*ResponseMembers, error) {
	ep := buildEndpoint(memberEndpointFmt, f.qp, f.qp.SortMembers)

	res, err := http.Get(ep)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := new(ResponseMembers)

	if err := json.Unmarshal(b, result); err != nil {
		return nil, err
	}

	return result, nil
}
