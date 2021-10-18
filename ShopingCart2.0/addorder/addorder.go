package addorder

import (
	"bufio"
	"dbconnection/test/order"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

func AddOrder(db *gorm.DB, id string, orders []*order.Order) []*order.Order {
	var itemName, itemDesc, tempnum string
	var quantity int64
	var isPaid bool
	var costPerUnit float64
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var flag uint64
		fmt.Println("Dear Customer: \n Enter 1 for order \n Enter 0 for back")
		for {
			var tempNum string
			if scanner.Scan() {
				tempNum = scanner.Text()
			}
			flag, err = strconv.ParseUint(tempNum, 10, 32)
			if err != nil {
				fmt.Println(errors.New("You have inserted invalid Value,Enter Numeric value"))
			} else {
				break
			}
		}
		if flag == 1 {
			fmt.Println("Enter the Item name which u want to buy:")
			if scanner.Scan() {
				itemName = scanner.Text()
			}
			fmt.Println("Enter the Item Discription:")
			if scanner.Scan() {
				itemDesc = scanner.Text()
			}
			for {
				fmt.Println("Enter the number of qunatity u want to buy:")
				if scanner.Scan() {
					tempnum = scanner.Text()
				}
				quantity, err = strconv.ParseInt(tempnum, 10, 32)
				if err != nil {
					fmt.Println(errors.New("You have inserted invalid quantity number"))
				} else {
					break
				}
			}
			for {
				fmt.Println("Enter cost of per unit of :", itemName)
				if scanner.Scan() {
					tempnum = scanner.Text()
				}
				costPerUnit, err = strconv.ParseFloat(tempnum, 32)
				if err != nil {
					fmt.Println(errors.New("You have inserted invalid cost float number"))
				} else {
					break
				}
			}
			for {
				fmt.Println("Enter true or false if u paid total cost")
				if scanner.Scan() {
					tempnum = scanner.Text()
				}
				isPaid, err = strconv.ParseBool(tempnum)
				if err != nil {
					fmt.Println(errors.New("You have inserted invalid value"))
				} else {
					break
				}
			}
			o := order.New(id, itemName, itemDesc, int(quantity), costPerUnit, &isPaid)
			if isPaid {
				orders = append(orders, o)
			}
			db.Create(o)
		} else if flag > 1 {
			fmt.Println("Enter valid choice number, Enter 1 or 0")
		} else {
			break
		}
	}
	return orders
}
