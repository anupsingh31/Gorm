package main

import (
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
	IsMale *bool
}

func main() {
	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	db.AutoMigrate(&Student4{})
	//fmt.Println("error is", db.Error)
	//db.DB().Query("SET SESSION SQL_MODE='ALLOW_INVALID_DATES")
	db.AutoMigrate(&Student4{})
	//db.DB().Query("SET SESSION SQL_MODE='ALLOW_INVALID_DATES")
	// db.Create(&Student4{FName: "Anmol", LName: "Patak", Age: 24, IsMale: true})
	// db.Create(&Student4{FName: "Saket", LName: "Mishra", Age: 22, IsMale: true})
	// db.Create(&Student4{FName: "Anupam", LName: "Singh", Age: 21, IsMale: true})
	// db.Create(&Student4{FName: "Ritika", LName: "Singh", Age: 21, IsMale: false})

	//db.DB().Query("SET SESSION SQL_MODE='ALLOW_INVALID_DATES")

	var student Student4
	// db.First(&student)
	// db.Delete(Student4{}, 3)
	// fmt.Println("First name ", student.FName, "last name : ", student.LName, "Age : ", student.Age)
	//db.Delete(Student4{}, []int{8, 9, 10})
	//db.Last(&Student4{})
	//fmt.Println("Last value ", , "last name : ", student.LName, "Age : ", student.Age)
	var male = false
	student = Student4{
		FName:  "Raj",
		LName:  "Kundra",
		Age:    28,
		IsMale: &male,
	}
	//db.Model(&student).Where("l_name = ?", "Singh").Update("FName", "Anupam")
	db.Debug().Model(&student).Where("l_name = ?", "Kundra").Update(student)

	//db.Model(&student).Where("age = 21").Update("age", 15)
	//db.Debug().Model(&student).Where("f_name = ?", "Sidesh").Update("IsMale", false)

	//db.Model(&student).Where("l_name = ?", "Patak").Updates(map[string]interface{}{"FName": "Sidesh", "LName": "kapoor", "Age": 26, "IsMale": false})
	//db.Model(Student4{}).Update("fname", "Abhi")
	// rows, err := db.DB().Query("SELECT f_name,age FROM student4")
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
