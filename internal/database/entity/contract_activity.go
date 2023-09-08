package entity

import (
	"time"
)

type ContractActivityType uint8

const (
	ContractActivityTypeNone ContractActivityType = iota
	ContractActivityTypeIntent
	ContractActivityTypePurchase
	ContractActivityTypeCancellation
)

type ContractActivity struct {
	ID         string               `db:"id"`
	ContractID string               `db:"contract_id"`
	Type       ContractActivityType `db:"type"`
	OccurredAt time.Time            `db:"occurred_at"`
	CreatedAt  time.Time            `db:"created_at"`
	UpdatedAt  time.Time            `db:"updated_at"`
}

type ContractIntent struct {
	ContractActivityID string    `db:"contract_activity_id"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

type ContractPurchase struct {
	ContractActivityID string    `db:"contract_activity_id"`
	Price              uint32    `db:"price"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

type ContractCancellationReason uint8

const (
	ContractCancellationReasonNone = iota
)

type ContractCancellation struct {
	ContractActivityID string                     `db:"contract_activity_id"`
	Reason             ContractCancellationReason `db:"reason"`
	CreatedAt          time.Time                  `db:"created_at"`
	UpdatedAt          time.Time                  `db:"updated_at"`
}
