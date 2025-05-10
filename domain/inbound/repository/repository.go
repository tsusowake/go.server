//go:generate mockgen -destination=mock/mock_repository.go github.com/tsusowake/go.server/domain/inbound/repository InboundRepository
package repository

type InboundRepository interface{}

type repository struct{}

var _ InboundRepository = (*repository)(nil)

func NewInboundRepository() InboundRepository {
	return &repository{}
}
