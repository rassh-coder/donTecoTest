package repository

import (
	"donTecoTest/pkg/models"
	"github.com/jackc/pgx/v5"
)

type IEmployee interface {
	FindByName(name string) (*models.Employee, error)
	GetList(limit, offset uint) (*[]models.Employee, error)
}

type Repository struct {
	Employee IEmployee
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{Employee: NewEmployeeRepository(db)}
}
