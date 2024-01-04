package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

type Database interface {
	InitDB(config Config) (*sql.DB, error)
	CloseDB() error
	Create(data interface{}) error
	Read(condition string) (interface{}, error)
	Update(data interface{}) error
	Delete(id int) error
}

type MySQLDB struct {
	db *sql.DB
}

// ...

func (m *MySQLDB) InitDB(config Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Database)
	var err error
	m.db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Verificar a conexão com o banco de dados
	if err = m.db.Ping(); err != nil {
		return nil, err
	}

	log.Printf("Conexão com o banco de dados estabelecida com sucesso para o banco: %s", config.Database)
	return m.db, nil
}

// ...

func (m *MySQLDB) CloseDB() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}

func (m *MySQLDB) Create(data interface{}) error {
	// Implementar a lógica de criação no MySQL
	return nil
}

func (m *MySQLDB) Read(condition string) (interface{}, error) {
	// Implementar a lógica de leitura no MySQL
	return nil, nil
}

func (m *MySQLDB) Update(data interface{}) error {
	// Implementar a lógica de atualização no MySQL
	return nil
}

func (m *MySQLDB) Delete(id int) error {
	// Implementar a lógica de exclusão no MySQL
	return nil
}
