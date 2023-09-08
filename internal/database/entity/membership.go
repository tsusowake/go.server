package entity

import (
	"time"
)

type MembershipRank uint8

const (
	MembershipRankNone MembershipRank = iota
	MembershipRankFree
	MembershipRankPremium
)

type Membership struct {
	UserID    string         `db:"user_id"`
	Rank      MembershipRank `db:"rank"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}
