package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Employees fields collection
type Employees struct {
	Model

	ID              int       `gorm:"column:employee_id;type:smallint;not null"`
	LastName        string    `gorm:"column:last_name;type:character;not null"`
	FistName        string    `gorm:"column:first_name;type:character;not null"`
	Title           string    `gorm:"column:title;type:character"`
	TitleOfCourtesy string    `gorm:"column:title_of_courtesy;type:character"`
	BirthDate       time.Time `gorm:"column:birth_date"`
	HireDate        time.Time `gorm:"column:hire_date"`
	Address         string    `gorm:"type:character"`
	City            string    `gorm:"type:character"`
	Region          string    `gorm:"type:character"`
	PostalCode      string    `gorm:"column: postal_code"`
	Country         string    `gorm:"type:character"`
	HomePhone       string    `gorm:"column:home_phone"`
	Extension       string    `gorm:"type:character"`
	Notes           string    `gorm:"type:text"`
	Reports         int64     `gorm:"column:reports_to;type:smallint"`
	Photo           string    `gorm:"type:character"`
}

// GetEmployees get a list of employees based on paging constraints
func GetEmployees() ([]*Employees, error) {
	var employees []*Employees
	err := db.Find(&employees).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return employees, nil
}
