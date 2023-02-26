package types

import (
	"go.olapie.com/utils"
	"testing"
)

func TestM_AddStruct(t *testing.T) {
	m := M{}
	m["id"] = 1
	m["name"] = "Smith"
	var foo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	foo.Name = "Mike"
	foo.Age = 19
	err := m.AddStruct(foo)
	utils.MustNotErrorT(t, err)
	utils.MustEqualT(t, 1, m.Int("id"))
	utils.MustEqualT(t, 19, m.Int("age"))
	utils.MustEqualT(t, "Mike", m.String("name"))
}
