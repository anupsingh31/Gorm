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
	db.AutoMigrate(&customer.Customer{})
	db.Create(c)
	db.AutoMigrate(&order.Order{})
	db.Model(&order.Order{}).AddForeignKey("customer_id", "customers(id)", "RESTRICT", "RESTRICT")
	for {
		var flag bool
		fmt.Println("Dear Customer: \n Enter 1 for order \n Enter 0 for checkout")
		fmt.Scanln(&flag)
		if flag {
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
			db.Create(o)
		} else {
			inVoice(*c, db)
			break
		}
	}
	// fmt.Println("Enter the Item name which u want to buy:")
	// fmt.Scanln(&itemName)
	// fmt.Println("Enter the Item Discription:")
	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	itemDesc = scanner.Text()
	// }
	// fmt.Println("Enter the number of qunatity u want to buy:")
	// fmt.Scanln(&quantity)
	// fmt.Println("Enter cost of per unit of  ", itemName)
	// fmt.Scanln(&costPerUnit)
	// fmt.Println("Enter true or false if u paid total cost")
	// fmt.Scanln(&isPaid)
	// o := order.New(c.ID, itemName, itemDesc, quantity, costPerUnit, &isPaid)
	// db.AutoMigrate(&order.Order{})
	// db.Create(o)
	// inVoice(*c, *o)
}

func inVoice(c customer.Customer, db *gorm.DB) {
	cl := exec.Command("cmd", "/c", "cls")
	cl.Stdout = os.Stdout
	cl.Run()
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Name    :    ", c.FName, c.LName)
	fmt.Println("Age     :    ", c.Age)
	fmt.Println("IsMale  :    ", *c.IsMale)
	fmt.Println("\n \n")
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Item name\t\t| Item Desc\t\t| Quantity\t\t| Cost/unit\t\t| ")
	var total float64 = 0
	rows, err := db.DB().Query("SELECT item_name,item_desc,quantity,cost_per_unit,is_paid FROM orders WHERE customer_id = ?", c.ID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var itemName, itemDesc string
		var quantity int
		var costPerUnit float64
		var isPaid bool
		err = rows.Scan(&itemName, &itemDesc, &quantity, &costPerUnit, &isPaid)
		if err != nil {
			panic(err)
		}
		fmt.Println(itemName, "\t| ", itemDesc, "\t|", quantity, "\t|", costPerUnit, "\t|")
		total += float64(quantity) * costPerUnit
	}
	fmt.Println("----------------------------------------------------------")
	fmt.Println("\t\t    \t\t     Total Cost|       ", total)
	fmt.Println("----------------------------------------------------------")

	if err != nil {
		panic(err)
	}
	// fmt.Println("Item Name                 |       ", o)
	// fmt.Println("Item Description          |       ", o.ItemDesc)
	// fmt.Println("Qantity                   |       ", o.Quantity)
	// fmt.Println("Per Unit Cost             |       ", o.CostPerUnit)
	// fmt.Println("IsPaid                    |       ", *o.IsPaid)
	// fmt.Println("\n \n \n")
	// fmt.Println("----------------------------------------------------------")
	// fmt.Println("                Total Cost|       ", o.CostPerUnit*float64(o.Quantity))
	// fmt.Println("----------------------------------------------------------")

}
