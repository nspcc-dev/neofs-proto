package eacl

import (
	"testing"

	v2acl "github.com/nspcc-dev/neofs-api-go/v2/acl"
	"github.com/stretchr/testify/require"
)

func newObjectFilter(match Match, key, val string) *Filter {
	return &Filter{
		from: HeaderFromObject,
		key: filterKey{
			str: key,
		},
		matcher: match,
		value:   staticStringer(val),
	}
}

func TestFilter(t *testing.T) {
	filter := newObjectFilter(MatchStringEqual, "some name", "200")

	v2 := filter.ToV2()
	require.NotNil(t, v2)
	require.Equal(t, v2acl.HeaderTypeObject, v2.GetHeaderType())
	require.EqualValues(t, v2acl.MatchTypeStringEqual, v2.GetMatchType())
	require.Equal(t, filter.Key(), v2.GetKey())
	require.Equal(t, filter.Value(), v2.GetValue())

	newFilter := NewFilterFromV2(v2)
	require.Equal(t, filter, newFilter)

	t.Run("from nil v2 filter", func(t *testing.T) {
		require.Equal(t, new(Filter), NewFilterFromV2(nil))
	})
}
