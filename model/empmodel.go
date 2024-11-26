package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)


type Emp struct{
	gorm.Model
	Id int  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null"`
	Design string `json:"email" gorm:"not null"`
	Salary int `json:"salary"`
}

func (e Emp) Validation() error{
	if strings.Trim(e.Name," ") == ""{
		return fmt.Errorf("Name should not be empty..")
	}

	if strings.Trim(e.Design," ") == ""{
		return fmt.Errorf("Design should not be empty..")
	}
	return nil
}
