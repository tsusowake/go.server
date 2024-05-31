package ulid

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateULID(t *testing.T) {
	t.Run("ULIDが生成できること", func(t *testing.T) {
		generator := NewULIDGenerator()
		t1 := time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC)
		got, err := generator.Generate(t1)
		assert.NotEmpty(t, got)
		assert.NoError(t, err)
	})

	t.Run("同一のタイムスタンプから異なるULIDが生成されること", func(t *testing.T) {
		generator := NewULIDGenerator()
		t1 := time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC)
		got1, err1 := generator.Generate(t1)
		got2, err2 := generator.Generate(t1)
		assert.NotEqual(t, got1, got2)
		assert.NoError(t, err1)
		assert.NoError(t, err2)
	})
}
