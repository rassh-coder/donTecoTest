package service

import (
	"donTecoTest/pkg/models"
	"donTecoTest/pkg/repository"
)

type EmployeeService struct {
	repos repository.IEmployee
}

func NewEmployeeService(repos repository.IEmployee) *EmployeeService {
	return &EmployeeService{repos: repos}
}

// Поиск работника по имени
func (s *EmployeeService) FindByName(name string) (*models.Employee, error) {
	emp, err := s.repos.FindByName(name)
	if err != nil {
		return nil, err
	}

	return emp, nil
}

func (s *EmployeeService) GetList(limit, offset uint) (*[]models.Employee, error) {
	emps, err := s.repos.GetList(limit, offset)
	if err != nil {
		return nil, err
	}

	return emps, nil
}
