package database

import (
	"fmt"
	"gorm_emp/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//connection string, automigration

type DBConfig struct{
	Host string
	Port string
	User string
	Password string
	DBName   string
}

var dbConfig = DBConfig{
	Host:     "127.0.0.1",
	Port:     "3306",
	User:     "root",
	Password: "abc123",
	DBName:   "store",
}



func InitDB() (*gorm.DB , error){
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	) 

    //db open 
    db, err := gorm.Open(mysql.Open(connectionString))
	if err != nil{
		return nil, fmt.Errorf("failed to connect to Mysql database : %v", err)
	}

    //db automation
	err = db.AutoMigrate(model.Emp{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate databse: %v", err)
	}

	return db, nil
}


