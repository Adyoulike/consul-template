package dependency

import (
	"sort"

	"github.com/hashicorp/consul/api"
)

// Dependency is an interface for a dependency that Consul Template is capable
// of watching.
type Dependency interface {
	Fetch(*api.Client, *api.QueryOptions) (interface{}, *api.QueryMeta, error)
	HashCode() string
	Display() string
}

// deepCopyAndSortTags deep copies the tags in the given string slice and then
// sorts and returns the copied result.
func deepCopyAndSortTags(tags []string) []string {
	newTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		newTags = append(newTags, tag)
	}
	sort.Strings(newTags)
	return newTags
}
