package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Students struct {
	gorm.Model
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
	db.AutoMigrate(&Students{})
	// db.Create(&Students{FName: "Saket", LName: "Mishra", Age: 22, Sex: "Male"})
	// db.Create(&Students{FName: "Ankit", LName: "Shah", Age: 22, Sex: "Male"})
	// db.Create(&Students{FName: "Abhi", LName: "Khtoria", Age: 25, Sex: "Male"})
	// db.Create(&Students{FName: "Anmol", LName: "Tilak", Age: 24, Sex: "Male"})
	var students Students
	db.Debug().Last(&students)
	fmt.Println("First Name :", students.FName, "Last Name:", students.LName, "Age : ", students.Age, "Sex : ", students.Sex)

	students = Students{}
	db.Debug().First(&students)
	fmt.Println("First Name :", students.FName, "Last Name:", students.LName, "Age : ", students.Age, "Sex : ", students.Sex)
	//var students1 Students

	db.Debug().Last(&students)
	///db.Model(&students).Update("Age", 23)
	//fmt.Println("First Name :", students.FName, "Last Name:", students.LName, "Age : ", students.Age, "Sex : ", students.Sex)
	//fmt.Println(db.Last(&students))
	//db.Unscoped().Where("age = 24").Delete(&Students{})
	db.Unscoped().Delete(&students)

	//rows, err := db.DB().Query("SELECT f_name,age FROM students")
	// rows:= db.First(&Students{})

	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var age int
	// 	var firstName string
	// 	err = rows.Scan(&firstName, &age)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(age, firstName)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
}
