//go:generate mockgen -destination=mock/mock_repository.go github.com/tsusowake/go.server/domain/outbound/repository OutboundRepository
package repository

type OutboundRepository interface{}

type repository struct{}

var _ OutboundRepository = (*repository)(nil)

func NewOutboundRepository() OutboundRepository {
	return &repository{}
}
