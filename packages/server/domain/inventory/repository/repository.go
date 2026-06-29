//go:generate mockgen -destination=mock/mock_repository.go github.com/tsusowake/go.server/domain/inventory/repository InventoryRepository
package repository

type InventoryRepository interface{}

type repository struct{}

var _ InventoryRepository = (*repository)(nil)

func NewInventoryRepository() InventoryRepository {
	return &repository{}
}
