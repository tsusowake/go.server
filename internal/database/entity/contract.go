package entity

import (
	"time"
)

type Plan struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Price     int64     `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Contract struct {
	ID        string    `db:"id"`
	UserID    string    `db:"uesr_id"`
	PlanID    string    `db:"plan_id"`
	StartedAt time.Time `db:"started_at"`
	EndedAt   time.Time `db:"ended_at"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// 契約請求

type ContractActivityType string

const (
	ContractActivityTypeNone = iota
	// キャンセル, 自動更新, ディスカウント, 購入
)

type ContractActivity struct {
	ID           string               `db:"id"`
	ContractID   string               `db:"contract_id"`
	ActivatedAt  time.Time            `db:"activated_at"`
	ActivityType ContractActivityType `db:"activity_type"`
	CreatedAt    time.Time            `db:"created_at"`
	UpdatedAt    time.Time            `db:"updated_at"`
}

type ContractCancellationReason uint16

const (
	ContractCancellationReasonNone = iota
)

type ContractCancellation struct {
	ContractActivityID string                     `db:"contract_activity_id"`
	Reason             ContractCancellationReason `db:"reason"`
	CanceledAt         time.Time                  `db:"canceled_at"`
	CreatedAt          time.Time                  `db:"created_at"`
	UpdatedAt          time.Time                  `db:"updated_at"`
}

type ContractAutoRenewal struct {
	ContractActivityID string    `db:"contract_activity_id"`
	RenewedAt          time.Time `db:"renewed_at"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

type ContractDiscountReason uint16

const (
	ContractDiscountReasonNone = iota
)

type ContractDiscount struct {
	ContractActivityID string                 `db:"contract_activity_id"`
	OriginalPrice      int64                  `db:"original_price"`
	DiscountAmount     int64                  `db:"discount_amount"`
	Reason             ContractDiscountReason `db:"reason"`
	GrossPrice         int64                  `db:"gross_price"`
	DiscountedAt       time.Time              `db:"discounted_at"`
	CreatedAt          time.Time              `db:"created_at"`
	UpdatedAt          time.Time              `db:"updated_at"`
}

// TODO ContractPurchase が Invoice とほぼ同じなので分ける必要があるのか!?
// サブスクリプションの契約は別アクティビティでやるべきか？
type ContractPurchase struct {
	ContractActivityID string    `db:"contract_activity_id"`
	InvoiceID          string    `db:"invoice_id"`
	Price              int64     `db:"price"`
	PurchasedAt        time.Time `db:"purchased_at"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

type InvoiceReason uint16

const (
	InvoiceReasonNone = iota
)

type Invoice struct {
	ID        string        `db:"id"`
	UserID    string        `db:"uesr_id"`
	BillsOn   time.Time     `db:"bills_on"`
	Price     int64         `db:"price"`
	Reason    InvoiceReason `db:"reason"`
	CreatedAt time.Time     `db:"created_at"`
	UpdatedAt time.Time     `db:"updated_at"`
}

type InvoiceActivityType string

const (
	InvoiceActivityTypeNone = iota
	// 請求完了, キャンセル, エラー
)

type InvoiceActivity struct {
	ID           string              `db:"id"`
	InvoiceID    string              `db:"invoice_id"`
	ActivatedAt  time.Time           `db:"activated_at"`
	ActivityType InvoiceActivityType `db:"activity_type"`
	CreatedAt    time.Time           `db:"created_at"`
	UpdatedAt    time.Time           `db:"updated_at"`
}

type InvoiceCancellationReason uint16

const (
	InvoiceCancellationReasonNone = iota
	//
)

type InvoiceCancellation struct {
	InvoiceActivityID string                    `db:"invoice_activity_id"`
	Reason            InvoiceCancellationReason `db:"reason"`
	CanceledAt        time.Time                 `db:"canceled_at"`
	CreatedAt         time.Time                 `db:"created_at"`
	UpdatedAt         time.Time                 `db:"updated_at"`
}

type InvoiceError struct {
	InvoiceActivityID string    `db:"invoice_activity_id"`
	ResponseBody      string    `db:"response_body"`
	ErroredAt         time.Time `db:"errored_at"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}

type InvoiceCompletion struct {
	InvoiceActivityID string    `db:"invoice_activity_id"`
	CompletedAt       time.Time `db:"completed_at"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
