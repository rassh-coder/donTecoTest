package migrationUp

import (
	"context"
	"donTecoTest/pkg/models"
	"encoding/csv"
	"fmt"
	"github.com/jackc/pgx/v5"
	"io"
	"log"
	"os"
	"strconv"
)

func Run(db *pgx.Conn) error {
	file, err := os.Open("./cmd/migration/up/data.csv")
	if err != nil {
		log.Printf("err: %s \n", err)
	}

	// Оставляю 50 для полей `employment`, `payment_system` как запаз для каких-то видоизменений подписей
	_, err = db.Exec(context.Background(),
		`CREATE TABLE IF NOT EXISTS employees (
			id bigserial primary key,
			name varchar(100),
			position text,
			department text,
			employment varchar(50), 
			payment_system varchar(50),
			typical_hours real,
			annual_salary real,
			hourly_rate real
			);`)

	if err != nil {
		log.Fatalf("can't create table:%s \n", err)
		return err
	}

	_, err = db.Exec(context.Background(), "CREATE INDEX IF NOT EXISTS name_idx on employees (name)")

	if err != nil {
		log.Printf("can't create index: %s \n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("can't close file: %s", err))
		}
	}(file)

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 8
	idx := 0
	for {
		record, e := reader.Read()
		// Пропуск первой строки с наименование столбцов
		if idx == 0 {
			idx++
			continue
		}

		if e == io.EOF {
			break
		}
		if e != nil {
			log.Printf("can't read: %s \n", e)
			return e
		}

		emp := models.Employee{
			Name:          record[0],
			Position:      record[1],
			Department:    record[2],
			Employment:    record[3],
			PaymentSystem: record[4],
		}

		//	Convert to float and fill float32 fields
		if record[5] != "" {
			f, err := strconv.ParseFloat(record[5], 32)
			if err != nil {
				log.Printf("Can't parse float: %s, emp name: %s \n", err, record[0])
			}
			emp.TypicalHours = float32(f)
		}

		if record[6] != "" {
			f, err := strconv.ParseFloat(record[6], 32)
			if err != nil {
				log.Printf("Can't parse float: %s, emp name: %s \n", err, record[0])
			}

			emp.AnnualSalary = float32(f)
		}

		if record[7] != "" {
			f, err := strconv.ParseFloat(record[7], 32)
			if err != nil {
				log.Printf("Can't parse float: %s, emp name: %s \n", err, record[0])
			}
			emp.HourlyRate = float32(f)
		}
		q := `INSERT INTO employees
			(name, position, department, employment, payment_system, typical_hours, annual_salary, hourly_rate)
			VALUES (@name, @position, @department, @employment, @paymentSystem, @typicalHours, @annualSalary, @hourlyRate)`
		args := pgx.NamedArgs{
			"name":          emp.Name,
			"position":      emp.Position,
			"department":    emp.Department,
			"employment":    emp.Employment,
			"paymentSystem": emp.PaymentSystem,
			"typicalHours":  emp.TypicalHours,
			"annualSalary":  emp.AnnualSalary,
			"hourlyRate":    emp.HourlyRate,
		}
		_, err = db.Exec(context.Background(), q, args)
		if err != nil {
			log.Printf("Can't insert row: %s, name: %s \n", err, emp.Name)
		}
	}

	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			log.Fatalf("can't close db connection")
		}
	}(db, context.Background())

	return nil
}
