package external

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	membersAPIEndpointFmt = EsaAPIOrigin +
		"/" +
		EsaAPIVersion +
		"/teams/%s/members?access_token=%s&sort=%s&order=%s&per_page=%d&page=%d"
)

type MembersAPIFetcher struct {
	qp       *APIQueryParams
	loopFlag *bool
}

type Member struct {
	Name string `json:"name"`
	// Myself 				bool   `json:"myself"`
	// ScreenName 		string `json:"screen_name"`
	// Icon 					string `json:"icon"`
	// Role 					string `json:"role"`
	// PostsCount 		int    `json:"posts_count"`
	// JoinedAt 			time.Time `json:"joined_at"`
	// LastAccessedAt time.Time `json:"last_accessed_at"`
}

type ResponseMembers struct {
	Members  []Member `json:"members"`
	NextPage int      `json:"next_page"`
	Page     int      `json:"page"`
	// PrevPage   int    `json:"prev_page"`
	// TotalCount int    `json:"total_count"`
	// PerPage    int    `json:"per_page"`
	// MaxPerPage int    `json:"max_per_page"`
}

func NewMembersAPIFetcher(qp *APIQueryParams) *MembersAPIFetcher {
	newTrue := true
	return &MembersAPIFetcher{qp: qp, loopFlag: &newTrue}
}

func (f *MembersAPIFetcher) Do() []Member {
	results := make([]Member, 0)

	for *f.loopFlag {
		result, err := f.do()
		if err != nil {
			panic(err)
		}

		results = append(results, result.Members...)

		if result.NextPage == 0 {
			*f.loopFlag = false
		}
		f.qp.Page = result.NextPage
	}

	return results
}

func (f *MembersAPIFetcher) do() (*ResponseMembers, error) {
	ep := buildAPIEndpoint(membersAPIEndpointFmt, f.qp, f.qp.SortMembers)

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
