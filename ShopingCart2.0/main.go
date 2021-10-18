package main

import (
	"bufio"
	"dbconnection/test/addorder"
	"dbconnection/test/cart"
	"dbconnection/test/customer"
	"dbconnection/test/login"
	"dbconnection/test/order"
	"dbconnection/test/previousorder"
	"dbconnection/test/registeruser"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	var c *customer.Customer
	var orders []*order.Order
	for {
		var flag = true
		fmt.Println("Input Value \n 1. For Register Type 1 \n 2. For Login Type 2")
		input := choiceNumber()
		switch input {
		case 1:
			c = registeruser.RegisterUserData()
			db.AutoMigrate(&customer.Customer{})
			db.Create(c)
			fmt.Println("You have successfully registred")
			c = login.UserLoginDetail(db)
		case 2:
			c = login.UserLoginDetail(db)

		default:
			flag = false
			fmt.Println("You have inserted wrong choice")
		}
		if flag {
			break
		}
	}
	for {
		var exit bool = false
		fmt.Println("Hello", c.FName, "\n Enter option \n 1. Add order \n 2. Previous Order \n 3. Cart \n 4. Checkout \n 5. Exit")
		input := choiceNumber()
		switch input {
		case 2:
			previousorder.PreviousOrder(db, c.ID)
		case 1:
			db.AutoMigrate(&order.Order{})
			db.Model(&order.Order{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "RESTRICT")
			orders = addorder.AddOrder(db, c.ID, orders)
		case 3:
			orders = cart.Cart(db, c.ID, orders)
		case 5:
			exit = true
			clearScreen()
		case 4:
			inVoice(*c, orders)
		default:
			fmt.Println("You have entered wrong choice")
		}
		if exit {
			break
		}
	}
}

func inVoice(c customer.Customer, orders []*order.Order) {
	clearScreen()
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Name    :    ", c.FName, c.LName)
	fmt.Println("Age     :    ", c.Age)
	fmt.Println("IsMale  :    ", *c.IsMale)
	fmt.Println("\n \n")
	fmt.Println("----------------------------------------------------------")
	var total float64 = 0
	fmt.Printf("%-10s %-20s %-20s %-20s %-20s\n", "Item No", "|Item name", "|Item Desc", "|Quantity", "|Cost/unit")
	for i := 0; i < len(orders); i++ {
		fmt.Printf("%-10d %-20s %-20s %-20d %-20f\n", i+1, orders[i].ItemName, orders[i].ItemDesc, orders[i].Quantity, orders[i].CostPerUnit)
		total += float64(orders[i].Quantity) * orders[i].CostPerUnit
	}
	if len(orders) == 0 {
		fmt.Println("You haven't any items on Cart")
	} else {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("\t\t    \t\t     Total Cost|       ", total)
		fmt.Println("----------------------------------------------------------")
	}
}
func clearScreen() {
	cl := exec.Command("cmd", "/c", "cls")
	cl.Stdout = os.Stdout
	cl.Run()
}

func choiceNumber() uint64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var tempNum string
		if scanner.Scan() {
			tempNum = scanner.Text()
		}
		num, err := strconv.ParseUint(tempNum, 10, 32)
		if err != nil {
			fmt.Println(errors.New("You have inserted invalid value, Enter Integer Number"))
		} else {
			return num
		}
	}
}
