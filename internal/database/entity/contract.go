package entity

import (
	"time"
)

type Platform uint8

const (
	PlatformNone Platform = iota
	PlatformWeb
	PlatformIOS
	PlatformAndroid
)

type Plan struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	ProductID string    `db:"product_id"`
	Platform  Platform  `db:"platform"`
	Price     uint32    `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ContractType uint16

const (
	ContractTypeNone ContractType = iota
	ContractTypeTrial
	ContractTypePremium
)

type Contract struct {
	ID        string       `db:"id"`
	UserID    string       `db:"user_id"`
	PlanID    string       `db:"plan_id"`
	Type      ContractType `db:"type"`
	StartedAt time.Time    `db:"started_at"`
	EndedAt   time.Time    `db:"ended_at"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
}

type PremiumPlanContract struct {
	ContractID     string    `db:"contract_id"`
	SubscriptionID string    `db:"subscription_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type TrialPlanContract struct {
	ContractID string    `db:"contract_id"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type ContractActivityType uint16

const (
	ContractActivityTypeNone ContractActivityType = iota
	ContractActivityTypeContractIntent
	ContractActivityTypeCancellation
	ContractActivityTypeDiscount
	ContractActivityTypeInvalidation
)

type ContractActivity struct {
	ID           string               `db:"id"`
	ContractID   string               `db:"contract_id"`
	ActivatedAt  time.Time            `db:"activated_at"`
	ActivityType ContractActivityType `db:"type"`
	CreatedAt    time.Time            `db:"created_at"`
	UpdatedAt    time.Time            `db:"updated_at"`
}

type ContractIntent struct {
	ContractActivityID string    `db:"contract_activity_id"`
	Price              uint32    `db:"price"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

type ContractCancellationReason uint16

const (
	ContractCancellationReasonNone = iota
)

type ContractCancellation struct {
	ContractActivityID string                     `db:"contract_activity_id"`
	Reason             ContractCancellationReason `db:"reason"`
	CreatedAt          time.Time                  `db:"created_at"`
	UpdatedAt          time.Time                  `db:"updated_at"`
}

type ContractDiscountReason uint16

const (
	ContractDiscountReasonNone = iota
)

type ContractDiscount struct {
	ContractActivityID string                 `db:"contract_activity_id"`
	OriginalPrice      uint32                 `db:"original_price"`
	DiscountAmount     uint32                 `db:"discount_amount"`
	Reason             ContractDiscountReason `db:"reason"`
	GrossPrice         uint32                 `db:"gross_price"`
	CreatedAt          time.Time              `db:"created_at"`
	UpdatedAt          time.Time              `db:"updated_at"`
}

//type InvoiceReason uint16
//
//const (
//	InvoiceReasonNone = iota
//)
//
//type Invoice struct {
//	ID        string        `db:"id"`
//	UserID    string        `db:"uesr_id"`
//	BillsOn   time.Time     `db:"bills_on"`
//	Price     int64         `db:"price"`
//	Reason    InvoiceReason `db:"reason"`
//	CreatedAt time.Time     `db:"created_at"`
//	UpdatedAt time.Time     `db:"updated_at"`
//}
//
//type InvoiceActivityType string
//
//const (
//	InvoiceActivityTypeNone = iota
//	// 請求完了, キャンセル, エラー
//)
//
//type InvoiceActivity struct {
//	ID           string              `db:"id"`
//	InvoiceID    string              `db:"invoice_id"`
//	ActivatedAt  time.Time           `db:"activated_at"`
//	ActivityType InvoiceActivityType `db:"activity_type"`
//	CreatedAt    time.Time           `db:"created_at"`
//	UpdatedAt    time.Time           `db:"updated_at"`
//}
//
//type InvoiceCancellationReason uint16
//
//const (
//	InvoiceCancellationReasonNone = iota
//	//
//)
//
//type InvoiceCancellation struct {
//	InvoiceActivityID string                    `db:"invoice_activity_id"`
//	Reason            InvoiceCancellationReason `db:"reason"`
//	CanceledAt        time.Time                 `db:"canceled_at"`
//	CreatedAt         time.Time                 `db:"created_at"`
//	UpdatedAt         time.Time                 `db:"updated_at"`
//}
//
//type InvoiceError struct {
//	InvoiceActivityID string    `db:"invoice_activity_id"`
//	ResponseBody      string    `db:"response_body"`
//	ErroredAt         time.Time `db:"errored_at"`
//	CreatedAt         time.Time `db:"created_at"`
//	UpdatedAt         time.Time `db:"updated_at"`
//}
//
//type InvoiceCompletion struct {
//	InvoiceActivityID string    `db:"invoice_activity_id"`
//	CompletedAt       time.Time `db:"completed_at"`
//	CreatedAt         time.Time `db:"created_at"`
//	UpdatedAt         time.Time `db:"updated_at"`
//}
