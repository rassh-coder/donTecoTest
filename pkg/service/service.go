package service

import (
	"donTecoTest/pkg/models"
	"donTecoTest/pkg/repository"
)

type IEmployee interface {
	FindByName(name string) (*models.Employee, error)
	GetList(limit, offset uint) (*[]models.Employee, error)
}

type Service struct {
	Employee IEmployee
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Employee: NewEmployeeService(repos.Employee),
	}
}
