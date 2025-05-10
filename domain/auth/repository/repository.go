//go:generate mockgen -destination=mock/mock_repository.go github.com/tsusowake/go.server/domain/auth/repository AuthRepository
package repository

type AuthRepository interface {
	FindUserByID(id string) (string, error)
	CreateUser() (string, error)
}

type authRepository struct{}

var _ AuthRepository = (*authRepository)(nil)

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) FindUserByID(id string) (string, error) {
	return "", nil
}

func (r *authRepository) CreateUser() (string, error) {
	return "", nil
}
