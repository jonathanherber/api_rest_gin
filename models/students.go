package models

import "gorm.io/gorm"

type Student struct { //struct for Students
	gorm.Model
	Name string `json:"name"` //formating to json
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
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
