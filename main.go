package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student4 struct {
	//gorm.Model
	FName string
	LName string
	Age   int
	Sex   string
}

func main() {
	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/student?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(&Student4{})
	// fmt.Println("error is", db.Error)
	db.AutoMigrate(&Student4{})
	// db.Create(&Student4{FName: "Anupam", LName: "Singh", Age: 21, Sex: "Male"})
	// db.Create(&Student4{FName: "Anmol", LName: "Patak", Age: 24, Sex: "Male"})
	// db.Create(&Student4{FName: "Saket", LName: "Mishra", Age: 22, Sex: "Male"})

	var student Student4
	db.First(&student)
	fmt.Println("First name ", student.FName, "last name : ", student.LName, "Age : ", student.Age)
	db.Delete(Student4{}, "id = ?", 3)
	db.Last(&student)
	fmt.Println("First name ", student.FName, "last name : ", student.LName, "Age : ", student.Age)

	db.Model(&student).Where("l_name = ?", "Singh").Update("FName", "Anupam")
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
