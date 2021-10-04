package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Student struct {
	//gorm.Model
	FName string
	LName string
	Age   int
	Sex   string
}

func main() {
	db, err := gorm.Open("mysql", "root:password@/swabhav?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(&Student{})
	// db.Create(&Student{FName: "Saket", LName: "Mishra", Age: 22, Sex: "Male"})
	// db.Create(&Student{FName: "Ankit", LName: "Shah", Age: 22, Sex: "Male"})
	// db.Create(&Student{FName: "Abhi", LName: "Khtoria", Age: 25, Sex: "Male"})
	// db.Create(&Student{FName: "Anmol", LName: "Tilak", Age: 24, Sex: "Male"})
	var students Student
	db.First(&students)
	// db.Model(&students).Update("Age", 23)
	fmt.Println("First Name :", students.FName, "Last Name:", students.LName, "Age : ", students.Age, "Sex : ", students.Sex)
	// db.Last(&students)
	db.Delete(&students, "f_name LIKE ?", "Anmol")
}
