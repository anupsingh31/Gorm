package main

import (
	"bufio"
	"dbconnection/test/customer"
	"dbconnection/test/order"
	"fmt"
	"log"
	"os"
	"os/exec"

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
	var fname, lname, itemName, itemDesc string
	var age, quantity int
	var ismale, isPaid bool
	var costPerUnit float64
	fmt.Println("Enter the customer First name:")
	fmt.Scanln(&fname)
	fmt.Println("Enter the customer Last name:")
	fmt.Scanln(&lname)
	fmt.Println("Enter the customer Age:")
	fmt.Scanln(&age)
	fmt.Println("Enter true or false customer is IsMale:")
	fmt.Scanln(&ismale)
	c := customer.New(fname, lname, age, &ismale)
	fmt.Println("Enter the Item name which u want to buy:")
	fmt.Scanln(&itemName)
	fmt.Println("Enter the Item Discription:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		itemDesc = scanner.Text()
	}
	fmt.Println("Enter the number of qunatity u want to buy:")
	fmt.Scanln(&quantity)
	fmt.Println("Enter cost of per unit of  ", itemName)
	fmt.Scanln(&costPerUnit)
	fmt.Println("Enter true or false if u paid total cost")
	fmt.Scanln(&isPaid)
	o := order.New(c.ID, itemName, itemDesc, quantity, costPerUnit, &isPaid)
	db.AutoMigrate(&customer.Customer{}, &order.Order{})
	db.Create(c)
	db.Create(o)
	inVoice(*c, *o)
}

func inVoice(c customer.Customer, o order.Order) {
	cl := exec.Command("cmd", "/c", "cls")
	cl.Stdout = os.Stdout
	cl.Run()
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Name    :    ", c.FName, c.LName)
	fmt.Println("Age     :    ", c.Age)
	fmt.Println("IsMale  :    ", *c.IsMale)
	fmt.Println("\n \n")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Item Name                 |       ", o.ItemName)
	fmt.Println("Item Description          |       ", o.ItemDesc)
	fmt.Println("Qantity                   |       ", o.Quantity)
	fmt.Println("Per Unit Cost             |       ", o.CostPerUnit)
	fmt.Println("IsPaid                    |       ", *o.IsPaid)
	fmt.Println("\n \n \n")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("                Total Cost|       ", o.CostPerUnit*float64(o.Quantity))
	fmt.Println("----------------------------------------------------------")

}
