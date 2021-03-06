package customer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	gorm.Model
	FName    string
	LName    string
	Age      int
	IsMale   *bool
	Email    string `gorm:"type:varchar(100);unique_index"`
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
