package cart

import (
	"bufio"
	"dbconnection/test/order"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/jinzhu/gorm"
)

func Cart(db *gorm.DB, id string, updatedOrder []*order.Order) []*order.Order {
	var tempNum string
	var exit = false
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%-10s %-20s %-20s %-20s %-20s\n", "Item No", "|Item name", "|Item Desc", "|Quantity", "|Cost/unit")
		var orders []*order.Order
		db.Where("customer_id = ? && is_paid = ?", id, false).Find(&orders)
		for i := 0; i < len(orders); i++ {
			fmt.Printf("%-10d %-20s %-20s %-20d %-20f\n", i+1, orders[i].ItemName, orders[i].ItemDesc, orders[i].Quantity, orders[i].CostPerUnit)
		}
		if len(orders) == 0 {
			fmt.Println("You haven't any items on Cart")
		}
		fmt.Println("Enter Input\n 1. For Delete Type 1 \n 2. For Update Quantity type 2 \n 3. For Update Is_order Type 3 \n 4. For back to Dashboard Type 4")
		if scanner.Scan() {
			tempNum = scanner.Text()
		}
		choice, err := strconv.ParseUint(tempNum, 10, 32)
		if err != nil {
			fmt.Println(errors.New("You have inserted invalid value"))
		} else {
			switch choice {
			case 1:
				num := itemNumber(len(orders))
				db.Where("id = ?", orders[num-1].ID).Delete(&orders)
				//clearScreen()
				fmt.Println("Successfully deleted item")
			case 2:
				num := itemNumber(len(orders))
				var quantity uint64
				var err error
				for {
					fmt.Println("enter the quantity value")
					if scanner.Scan() {
						tempNum = scanner.Text()
					}
					quantity, err = strconv.ParseUint(tempNum, 10, 32)
					if err != nil {
						fmt.Println(errors.New("You have inserted invalid value"))
					} else {
						break
					}
				}
				db.Model(&orders).Where("id = ?", orders[num-1].ID).Update("quantity", quantity)
				//clearScreen()
				fmt.Println("Succefully updated your quantity")
			case 3:
				num := itemNumber(len(orders))
				updatedOrder = append(updatedOrder, orders[num-1])
				db.Model(&orders).Where("id = ?", orders[num-1].ID).Update("is_paid", true)
				//clearScreen()
				fmt.Println("you have paid amount of that Item Value")
			case 4:
				exit = true
				clearScreen()
			default:
				clearScreen()
				fmt.Println("you have inserted wrong choice")
			}
			if exit {
				break
			}
		}
	}
	return updatedOrder
}
func itemNumber(size int) uint64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var tempNum string
		fmt.Println("For update item enter the Item Number")
		if scanner.Scan() {
			tempNum = scanner.Text()
		}
		num, err := strconv.ParseUint(tempNum, 10, 32)
		if err != nil {
			fmt.Println(errors.New("You have inserted invalid value"))
		} else if num > 0 && int(num) <= size {
			return num
		} else {
			fmt.Println("You have entered invalid Item No")
		}
	}
}

func clearScreen() {
	cl := exec.Command("cmd", "/c", "cls")
	cl.Stdout = os.Stdout
	cl.Run()
}
