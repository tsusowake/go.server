package slice

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFirst(t *testing.T) {
	type Element struct {
		ID   string
		Name string
	}

	t.Run("first", func(t *testing.T) {
		els := []Element{
			{ID: "id.1", Name: "name.1"},
			{ID: "id.2", Name: "name.2"},
			{ID: "id.3", Name: "name.3"},
		}
		got, err := First(els, func(v Element) bool {
			return v.ID == "id.2"
		})
		want := Element{ID: "id.2", Name: "name.2"}
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
	t.Run("empty", func(t *testing.T) {
		var els []Element
		got, err := First(els, func(v Element) bool {
			return v.ID == "id.2"
		})
		want := Element{ID: "", Name: ""}
		require.Error(t, err)
		require.Equal(t, want, got)
	})
	t.Run("missing", func(t *testing.T) {
		els := []Element{{ID: "id.1", Name: "name.1"}}
		got, err := First(els, func(v Element) bool {
			return v.ID == "id.999"
		})
		want := Element{ID: "", Name: ""}
		require.Error(t, err)
		require.Equal(t, want, got)
	})
}

func TestSelectString(t *testing.T) {
	// TODO
}

func TestWhere(t *testing.T) {
	// TODO
}
