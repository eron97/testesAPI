// services/user_service.go
package services

import (
	"fmt"

	"github.com/eron97/testesAPI/database"
)

// User representa a estrutura do usuário (exemplo básico)
type User struct {
	ID    int
	Name  string
	Email string
}

type UserService struct {
	DB database.Database
}

// CreateUser cria um novo usuário
func (s *UserService) CreateUser(user *User) error {
	// Implementação da lógica para criar um usuário
	return s.DB.Create(user)
}

// GetUserByID obtém um usuário pelo ID
func (s *UserService) GetUserByID(userID int) (*User, error) {
	// Implementação da lógica para obter um usuário pelo ID
	data, err := s.DB.Read(fmt.Sprintf("WHERE id = %d", userID))
	if err != nil {
		return nil, err
	}

	user, ok := data.(*User)
	if !ok {
		return nil, fmt.Errorf("Falha ao converter os dados do usuário")
	}

	return user, nil
}

// UpdateUser atualiza as informações de um usuário
func (s *UserService) UpdateUser(user *User) error {
	// Implementação da lógica para atualizar um usuário
	return s.DB.Update(user)
}

// DeleteUser exclui um usuário pelo ID
func (s *UserService) DeleteUser(userID int) error {
	// Implementação da lógica para excluir um usuário
	return s.DB.Delete(userID)
}

// ... outras funções de serviço
