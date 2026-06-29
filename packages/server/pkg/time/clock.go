package time

import "time"

type Clocker interface {
	Now() time.Time
}
type clocker struct{}

func (*clocker) Now() time.Time {
	return time.Now()
}

func NewClocker() Clocker {
	return &clocker{}
}
