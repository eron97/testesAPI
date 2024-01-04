// services/user_service.go
package services

import (
	"github.com/eron97/testesAPI/config/database"
	"github.com/eron97/testesAPI/config/models"
)

type UserServiceInterface interface {
	GetUsersWithProcessing(condition string) ([]models.User, error)
	// Adicione outras funções de serviço conforme necessário
}

// UserService é a implementação real do serviço de usuário
type UserService struct {
	DB database.Database
}

// GetUsersWithProcessing obtém usuários com algum processamento adicional
func (s *UserService) GetUsersWithProcessing(condition string) ([]models.User, error) {
	// Obter usuários do banco de dados
	data, err := s.DB.Read(condition)
	if err != nil {
		return nil, err
	}

	return data, nil
}
