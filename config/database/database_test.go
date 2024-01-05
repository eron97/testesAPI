package database_test

import (
	"errors"
	"testing"

	mocks "github.com/eron97/testesAPI/config/database/gomock"
	"github.com/eron97/testesAPI/config/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMySQLDB_Read(t *testing.T) {
	// Cria controlador mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Cria instância mock do banco de dados
	mockDB := mocks.NewMockDatabase(ctrl)

	// Definindo o comportamento esperado para o método Read
	expectedUsers := []models.User{{ID: 1, Username: "novo_usuario"}}
	mockDB.EXPECT().Read(gomock.Eq("SELECT id, username FROM users")).Return(expectedUsers, nil)

	users, err := mockDB.Read("SELECT id, username FROM users")
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)

	// Test for non-existent username:
	expectedError := errors.New("username not found")
	mockDB.EXPECT().Read(gomock.Eq("SELECT id, username FROM users")).Return(nil, expectedError)

	users, err = mockDB.Read("SELECT id, username FROM users")
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, users)

}
