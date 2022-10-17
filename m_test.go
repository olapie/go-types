package types_test

import (
	"testing"

	"code.olapie.com/types"
	"github.com/stretchr/testify/require"
)

func TestM_AddStruct(t *testing.T) {
	m := types.M{}
	m["id"] = 1
	m["name"] = "Smith"
	var foo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	foo.Name = "Mike"
	foo.Age = 19
	err := m.AddStruct(foo)
	require.NoError(t, err)
	require.Equal(t, 1, m.Int("id"))
	require.Equal(t, 19, m.Int("age"))
	require.Equal(t, "Mike", m.String("name"))
}
