package models

type Student struct { //struct for Students
	Name string `json:"name"` //formating to json
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Students []Student //slice of Student
