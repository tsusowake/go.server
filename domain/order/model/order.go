package model

// 合計金額
// 消費税
type Order struct {
	ID string
}

type OrderItem struct {
	ID      string
	OrderID string

	Product  Product
	Quantity uint16
	// ビジネスルール
}

// 消費税区分: 標準税率 or 軽減税率
// 酒税, たばこ税, ...
type Product struct{}

type ConsumptionTax struct{}
