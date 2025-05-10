//go:generate mockgen -destination=mock/mock_repository.go github.com/tsusowake/go.server/domain/order/repository OrderRepository
package repository

type OrderRepository interface {
	CreateOrder() (string, error)
	FindOrderByID(id string) (string, error)
}

type orderRepository struct{}

var _ OrderRepository = (*orderRepository)(nil)

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (r *orderRepository) CreateOrder() (string, error) {
	return "", nil
}

func (r *orderRepository) FindOrderByID(id string) (string, error) {
	return "", nil
}
