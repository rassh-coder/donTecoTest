package repository

import (
	"context"
	"donTecoTest/pkg/models"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

type EmployeeRepository struct {
	db *pgx.Conn
}

func NewEmployeeRepository(db *pgx.Conn) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) FindByName(name string) (*models.Employee, error) {
	emp := models.Employee{}
	q := `SELECT * FROM employees WHERE lower(name) = lower(@name)`
	args := pgx.NamedArgs{
		"name": name,
	}
	row := r.db.QueryRow(context.Background(), q, args)

	if err := row.Scan(
		&emp.Id, &emp.Name, &emp.Department,
		&emp.Employment, &emp.PaymentSystem, &emp.TypicalHours,
		&emp.AnnualSalary, &emp.HourlyRate,
	); err != nil {
		// Поиск похожих имен (на случай если нужен не жесткий поиск)
		if errors.Is(err, pgx.ErrNoRows) {
			q = `SELECT * FROM employees WHERE lower(name) like lower(@name)`
			args = pgx.NamedArgs{
				"name": fmt.Sprintf("%%%s%%", name),
			}
			row = r.db.QueryRow(context.Background(), q, args)

			if err = row.Scan(
				&emp.Id, &emp.Name, &emp.Position, &emp.Department,
				&emp.Employment, &emp.PaymentSystem, &emp.TypicalHours,
				&emp.AnnualSalary, &emp.HourlyRate,
			); err != nil {
				return nil, err
			}

			return &emp, nil
		}
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepository) GetList(limit, offset uint) (*[]models.Employee, error) {
	var emps []models.Employee
	var q string
	if limit > 0 {
		q = `SELECT * FROM employees ORDER BY id ASC OFFSET @offset LIMIT @limit`
	} else {
		q = `SELECT * FROM employees ORDER BY id ASC`
	}

	args := pgx.NamedArgs{
		"limit":  limit,
		"offset": offset,
	}

	rows, err := r.db.Query(context.Background(), q, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(&emp.Id, &emp.Name, &emp.Position,
			&emp.Department, &emp.Employment, &emp.PaymentSystem, &emp.TypicalHours, &emp.AnnualSalary, &emp.HourlyRate)
		if err != nil {
			log.Printf("Can't scan row: %s", err)
			continue
		}
		emps = append(emps, emp)
	}

	defer rows.Close()

	return &emps, nil
}
