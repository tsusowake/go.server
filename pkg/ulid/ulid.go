package ulid

import (
	"crypto/rand"
	"time"

	"github.com/morikuni/failure/v2"
	"github.com/oklog/ulid/v2"

	"github.com/tsusowake/go.server/pkg/errutil"
)

type ULIDGenerator interface {
	Generate(t time.Time) (string, error)
}
type ulidGenerator struct{}

func NewULIDGenerator() ULIDGenerator {
	return &ulidGenerator{}
}

func (_ *ulidGenerator) Generate(t time.Time) (string, error) {
	entropy := ulid.Monotonic(rand.Reader, 0)

	id, err := ulid.New(ulid.Timestamp(t.UTC()), entropy)
	if err != nil {
		return "", failure.New(
			errutil.ErrorCodeInternal,
			failure.Message("failed to generate ulid"),
		)
	}
	return id.String(), nil
}
