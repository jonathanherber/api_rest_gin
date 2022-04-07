package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct { //struct for Students
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf"  validate:"len=11, regexp=^[0-9]*$"`
	RG   string `json:"rg"   validate:"len=9, regexp=^[0-9]*$"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}

//var Students []Student //slice of Student

/*
//gorm.Model in the structure builds a db table with:
type Student struct {
ID unit `gorm:"primaryKey"`
CreatedTime time.Time
UpdatedTime time.Time
DeletedAt time.Time
DeletedAt gorm.DeletedAt `gorm:"index"`
}
//its called a embedded struct
*/
