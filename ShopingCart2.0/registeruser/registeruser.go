package registeruser

import (
	"bufio"
	"dbconnection/test/customer"
	"errors"
	"fmt"
	"net/mail"
	"os"
	"strconv"
)

func RegisterUserData() *customer.Customer {
	var fname, lname, email, tempnum string
	var age int64
	var ismale bool
	var password string
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the customer First name:")
	if scanner.Scan() {
		fname = scanner.Text()
	}
	fmt.Println("Enter the customer Last name:")
	if scanner.Scan() {
		lname = scanner.Text()
	}
	for {
		fmt.Println("Enter the customer Age:")
		if scanner.Scan() {
			tempnum = scanner.Text()
		}
		age, err = strconv.ParseInt(tempnum, 10, 64)
		if err != nil {
			fmt.Println(errors.New("You have inserted invalid age"))
		} else {
			break
		}
	}
	for {
		fmt.Println("Enter true or false customer is IsMale:")
		if scanner.Scan() {
			tempnum = scanner.Text()
		}
		ismale, err = strconv.ParseBool(tempnum)
		if err != nil {
			fmt.Println(errors.New("You have inserted invalid value"))
		} else {
			break
		}
	}
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
	fmt.Println("Enter the customer password:")
	fmt.Scanln(&password)
	c := customer.New(fname, lname, int(age), &ismale, email, password)
	return c
}
