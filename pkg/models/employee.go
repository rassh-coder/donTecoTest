package models

type Employee struct {
	Id            uint    `json:"id,omitempty"`
	Name          string  `json:"name,omitempty" json:"name,omitempty"`
	Position      string  `json:"position" json:"position,omitempty"`
	Department    string  `json:"department" json:"department,omitempty"`
	Employment    string  `json:"employment" json:"employment,omitempty"`
	PaymentSystem string  `json:"paymentSystem" json:"paymentSystem,omitempty"`
	TypicalHours  float32 `json:"typicalHours" json:"typicalHours,omitempty"`
	AnnualSalary  float32 `json:"annualSalary" json:"annualSalary,omitempty"`
	HourlyRate    float32 `json:"hourlyRate" json:"hourlyRate,omitempty"`
}

type EmployeeInputFields struct {
	Name string `json:"name,require"`
}
