//go:generate mockgen -destination=mock/mock_repository.go github.com/tsusowake/go.server/domain/purchase/repository PurchaseRepository
package repository

type PurchaseRepository interface{}

type repository struct{}

var _ PurchaseRepository = (*repository)(nil)

func NewPurchaseRepository() PurchaseRepository {
	return &repository{}
}
