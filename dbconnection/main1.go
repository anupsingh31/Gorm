package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student4 struct {
	gorm.Model
	FName  string
	LName  string
	Age    int
	IsMale bool
}

func main() {
	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	//db.AutoMigrate(&Student4{})
	// fmt.Println("error is", db.Error)
	//db.DB().Query("SET SESSION SQL_MODE='ALLOW_INVALID_DATES")
	//db.AutoMigrate(&Student4{})
	// db.DB().Query("SET SESSION SQL_MODE='ALLOW_INVALID_DATES")
	// db.Create(&Student4{FName: "Anmol", LName: "Patak", Age: 24, IsMale: true})
	// db.Create(&Student4{FName: "Saket", LName: "Mishra", Age: 22, IsMale: true})
	// db.Create(&Student4{FName: "Anupam", LName: "Singh", Age: 21, IsMale: true})
	//db.DB().Query("SET SESSION SQL_MODE='ALLOW_INVALID_DATES")

	var student Student4
	db.First(&student)
	fmt.Println("First name ", student.FName, "last name : ", student.LName, "Age : ", student.Age)
	db.Delete(Student4{}, "id = ?", 3)
	db.Last(&student)
	fmt.Println("First name ", student.FName, "last name : ", student.LName, "Age : ", student.Age)

	db.Model(&student).Where("l_name = ?", "Singh").Update("FName", "Anupam")
	db.Model(&student).Where("l_name = ?", "Patak").Updates(map[string]interface{}{"FName": "Sidesh", "LName": "kapoor", "Age": 26, "IsMale": false})
	//db.Model(Student4{}).Update("fname", "Abhi")
	rows, err := db.DB().Query("SELECT f_name,age FROM student4")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var age int
		var firstName string
		err = rows.Scan(&firstName, &age)
		if err != nil {
			panic(err)
		}
		fmt.Println(age, firstName)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
