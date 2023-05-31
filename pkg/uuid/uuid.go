package uuid

import (
	"encoding/base32"
	"github.com/google/uuid"
)

var (
	// see https://datatracker.ietf.org/doc/html/rfc4648#section-7
	encoding = base32.StdEncoding.WithPadding(base32.NoPadding)
)

func NewURLSafeString() (string, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return encoding.EncodeToString(v[:]), nil
}
