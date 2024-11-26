package service 

import (
	"fmt"
	"gorm_emp/model"
	"gorm.io/gorm"
)

func CreateService(emp model.Emp, db *gorm.DB) (model.Emp, error){
	valErr:= emp.Validation()
    if valErr != nil{
		return emp, valErr
	}

	result:= db.Create(&emp)
	if result.Error != nil{
		return emp, fmt.Errorf("User can not be created...")
	}

	return emp, nil
}

func GetService(db *gorm.DB)([]model.Emp, error){
		//make a list to get all elements.
		var emp []model.Emp
		result:= db.Find(&emp)
		if result.Error != nil{
			return emp, fmt.Errorf("Can not fetch emp from table...")
		}
		return emp, nil
	}


func GetbyIdService(id int, db *gorm.DB)(model.Emp, error){
	var emp model.Emp
	result:= db.First(&emp, id)
	if result.Error != nil{
		return emp, fmt.Errorf("Emp not found...")
	}
	return emp, nil
}


func UpdateIdService(id int, emp model.Emp, db *gorm.DB)(model.Emp, error){
    var updatedEmp model.Emp
	result:= db.First(&updatedEmp, id)
	if result.Error != nil{
		return updatedEmp, fmt.Errorf("Emp not found...")
	}

	result= db.Model(&updatedEmp).Updates(emp) 
	if result.Error != nil{
		return updatedEmp, fmt.Errorf("Emp updation failed...")
	}
    // fmt.Println(emp)
	// fmt.Println(updatedEmp)
	return updatedEmp, nil
}

func DeleteService(id int, db *gorm.DB) error{
	result:= db.Delete(&model.Emp{}, id)
	if result.Error != nil{
		return fmt.Errorf("Emp Delete failed...")
	}
	if result.RowsAffected == 0{
		return fmt.Errorf("Emp not found...")
	}
	return nil

}