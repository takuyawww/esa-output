package csv_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/takuyawww/esa-output/src/external"
	"github.com/takuyawww/esa-output/src/output/csv"
)

var mockMembers = []external.Member{
	{
		Name: "esako",
	},
	{
		Name: "esao",
	},
	{
		Name: "esami",
	},
}

var mockPosts = []external.Post{
	{
		Number:   1,
		Name:     "post1",
		Category: "first",
		CreatedBy: external.PostCreatedBy{
			Name: mockMembers[0].Name,
		},
		CreatedAt: time.Now(),
		UpdatedBy: external.PostUpdatedBy{
			Name: mockMembers[1].Name,
		},
		UpdatedAt: time.Now(),
		Wip:       true,
	},
	{
		Number:   2,
		Name:     "post2",
		Category: "first/second",
		CreatedBy: external.PostCreatedBy{
			Name: "not exist",
		},
		CreatedAt: time.Now(),
		UpdatedBy: external.PostUpdatedBy{
			Name: mockMembers[2].Name,
		},
		UpdatedAt: time.Now(),
		Wip:       false,
	},
	{
		Number:   3,
		Name:     "post3",
		Category: "first/second/third",
		CreatedBy: external.PostCreatedBy{
			Name: mockMembers[1].Name,
		},
		CreatedAt: time.Now(),
		UpdatedBy: external.PostUpdatedBy{
			Name: mockMembers[2].Name,
		},
		UpdatedAt: time.Now(),
		Wip:       true,
	},
	{
		Number:   4,
		Name:     "post4",
		Category: "first/second/third/forth",
		CreatedBy: external.PostCreatedBy{
			Name: "not exist",
		},
		CreatedAt: time.Now(),
		UpdatedBy: external.PostUpdatedBy{
			Name: mockMembers[1].Name,
		},
		UpdatedAt: time.Now(),
		Wip:       false,
	},
	{
		Number:   5,
		Name:     "post5",
		Category: "first/second/third/forth/fifth",
		CreatedBy: external.PostCreatedBy{
			Name: mockMembers[2].Name,
		},
		CreatedAt: time.Now(),
		UpdatedBy: external.PostUpdatedBy{
			Name: mockMembers[0].Name,
		},
		UpdatedAt: time.Now(),
		Wip:       true,
	},
}

func TestCsvToString(t *testing.T) {
	t.Run("test csv to string", func(t *testing.T) {
		csv := csv.New(mockMembers, mockPosts)
		fmt.Println(csv.String())
		// WIP
	})
}
