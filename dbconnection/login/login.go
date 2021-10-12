package login

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func UserLoginDetail(db *gorm.DB) uint {
	fmt.Println("Enter details for Login")
	var email string
	var flag bool = false
	var password string
	var id int
	fmt.Println("Enter the Email ")
	fmt.Scanln(&email)
	fmt.Println("Enter the password")
	fmt.Scanln(&password)
	rows, err := db.DB().Query("SELECT id,email,password FROM customers")
	if err != nil {
		panic(err)
	}
	fmt.Println(rows.Columns())
	defer rows.Close()
	for rows.Next() {
		var cEmail string
		var cPassword string
		err = rows.Scan(&id, &cEmail, &cPassword)
		fmt.Println(cEmail, cPassword)
		if err != nil {
			panic(err)
		}
		if cEmail == email && cPassword == password {
			flag = true
			rows.Close()
		}
	}
	if flag {
		fmt.Println("Succesfully login")
		return uint(id)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return 0
}
