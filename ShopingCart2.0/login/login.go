package login

import (
	"bufio"
	"dbconnection/test/customer"
	"fmt"
	"net/mail"
	"os"

	"github.com/jinzhu/gorm"
)

func UserLoginDetail(db *gorm.DB) *customer.Customer {
	fmt.Println("Enter details for Login")
	var email string
	var password string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		for {
			fmt.Println("Enter the Email ")
			if scanner.Scan() {
				email = scanner.Text()
			}
			_, err := mail.ParseAddress(email)
			if err == nil {
				break
			}
			fmt.Println(err)
		}
		fmt.Println("Enter the password")
		fmt.Scanln(&password)
		var customer1 []*customer.Customer
		db.Where("email = ? AND password = ?", email, password).Find(&customer1)
		if len(customer1) != 0 {
			fmt.Println("Hello ", customer1[0].FName, " is Successfully Login")
			return customer1[0]
		} else {
			fmt.Println("You have entered wrong email or password")
		}
	}
}
