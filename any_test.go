package types_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"code.olapie.com/types"
	"github.com/stretchr/testify/require"
)

func jsonString(i any) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func nextImage() *types.Image {
	return &types.Image{
		Url:    "https://www.image.com/" + fmt.Sprint(time.Now().Unix()),
		Width:  rand.Int31(),
		Height: rand.Int31(),
		Format: "png",
	}
}

func nextVideo() *types.Video {
	return &types.Video{
		Url:      "http://www.video.com/" + fmt.Sprint(time.Now().Unix()),
		Format:   "rmvb",
		Duration: rand.Int31(),
		Size:     rand.Int31(),
		Image:    nextImage(),
	}
}

func TestAny(t *testing.T) {
	t.Run("AliasType", func(t *testing.T) {
		type ID int
		types.RegisterAnyType(ID(0))

		v := types.NewAny(ID(10))
		b, err := json.Marshal(v)
		require.NoError(t, err)

		var vv *types.Any
		err = json.Unmarshal(b, &vv)
		require.NoError(t, err)

		require.Equal(t, jsonString(v), jsonString(vv))
	})

	t.Run("String", func(t *testing.T) {
		v := types.NewAny("hello")
		b, err := json.Marshal(v)
		require.NoError(t, err)

		var vv *types.Any
		err = json.Unmarshal(b, &vv)
		require.NoError(t, err)

		require.Equal(t, jsonString(v), jsonString(vv))
	})

	t.Run("Struct", func(t *testing.T) {
		v := types.NewAny(nextVideo())
		b, err := json.Marshal(v)
		require.NoError(t, err)

		var vv *types.Any
		err = json.Unmarshal(b, &vv)
		require.NoError(t, err)

		require.Equal(t, jsonString(v), jsonString(vv))
	})

	t.Run("Array", func(t *testing.T) {
		var l []*types.Any
		l = append(l, types.NewAny("hello"))
		l = append(l, types.NewAny(nextImage()))
		l = append(l, types.NewAny(nextVideo()))
		b, err := json.Marshal(l)
		require.NoError(t, err)

		var ll []*types.Any
		err = json.Unmarshal(b, &ll)
		require.NoError(t, err)

		require.Equal(t, jsonString(l), jsonString(ll))
	})
}
