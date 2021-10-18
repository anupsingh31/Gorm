package customer

import (
	"dbconnection/test/model"
)

type Customer struct {
	model.Model
	FName    string
	LName    string
	Age      int
	IsMale   *bool
	Email    string `gorm:"uniqueIndex"`
	Password string
}

func New(cFName, cLName string, cAge int, cIsMale *bool, email string, password string) *Customer {
	return &Customer{
		FName:    cFName,
		LName:    cLName,
		Age:      cAge,
		IsMale:   cIsMale,
		Email:    email,
		Password: password,
	}
}

// func (c *Customer) InsertCustomerTable() {
// 	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
// 	defer db.Close()
// 	if err != nil {
// 		log.Println("Connection Failed to Open")
// 	} else {
// 		log.Println("Connection Established")
// 	}
// 	db.AutoMigrate(&Customer{})
// 	db.Create(c)
// }
// func (c *Customer) GetCustomerID() uint {
// 	return c.ID
// }
