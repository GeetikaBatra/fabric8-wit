package application

import (
	"github.com/fabric8-services/fabric8-wit/workitem"
	"github.com/fabric8-services/fabric8-wit/workitem/link"

	"context"
)

// SearchRepository encapsulates searching of woritems,users,etc
type SearchRepository interface {
	SearchFullText(ctx context.Context, searchStr string, start *int, length *int, spaceID *string) ([]workitem.WorkItem, int, error)
	Filter(ctx context.Context, filterStr string, parentExists *bool, start *int, length *int) ([]workitem.WorkItem, int, link.AncestorList, link.WorkItemLinkList, error)
}
