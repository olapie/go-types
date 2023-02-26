package types

import (
	"encoding/json"
	"go.olapie.com/utils"
	"sort"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	s1 := NewSet[int](10)
	a0 := []int{1, 2, 3, 5, 9}
	for _, v := range a0 {
		s1.Add(v)
	}
	d1, err := s1.MarshalJSON()
	utils.MustNotErrorT(t, err)
	var s2 *Set[int]
	err = json.Unmarshal(d1, &s2)
	utils.MustNotErrorT(t, err)
	a1 := s1.Slice()
	a2 := s2.Slice()
	sort.IntSlice(a1).Sort()
	sort.IntSlice(a2).Sort()
	utils.MustEqualT(t, a0, a1)
	utils.MustEqualT(t, a1, a2)
}
