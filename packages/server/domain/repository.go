//go:generate mockgen -source=repository.go -destination=mock/mock_repository.go github.com/tsusowake/go.server/domain Repository
package domain

import (
	auth "github.com/tsusowake/go.server/domain/auth/repository"
	inbound "github.com/tsusowake/go.server/domain/inbound/repository"
	inventory "github.com/tsusowake/go.server/domain/inventory/repository"
	order "github.com/tsusowake/go.server/domain/order/repository"
	outbound "github.com/tsusowake/go.server/domain/outbound/repository"
	purchase "github.com/tsusowake/go.server/domain/purchase/repository"
)

type Repository interface {
	Auth() auth.AuthRepository
	Inbound() inbound.InboundRepository
	Inventory() inventory.InventoryRepository
	Order() order.OrderRepository
	Outbound() outbound.OutboundRepository
	Purchase() purchase.PurchaseRepository
}

type repository struct {
	auth      auth.AuthRepository
	inbound   inbound.InboundRepository
	inventory inventory.InventoryRepository
	order     order.OrderRepository
	outbound  outbound.OutboundRepository
	purchase  purchase.PurchaseRepository
}

func NewRepository() Repository {
	return &repository{
		auth:      auth.NewAuthRepository(),
		inbound:   inbound.NewInboundRepository(),
		inventory: inventory.NewInventoryRepository(),
		order:     order.NewOrderRepository(),
		outbound:  outbound.NewOutboundRepository(),
		purchase:  purchase.NewPurchaseRepository(),
	}
}

func (r *repository) Auth() auth.AuthRepository                { return r.auth }
func (r *repository) Inbound() inbound.InboundRepository       { return r.inbound }
func (r *repository) Inventory() inventory.InventoryRepository { return r.inventory }
func (r *repository) Order() order.OrderRepository             { return r.order }
func (r *repository) Outbound() outbound.OutboundRepository    { return r.outbound }
func (r *repository) Purchase() purchase.PurchaseRepository    { return r.purchase }
