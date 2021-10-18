package previousorder

import (
	"bufio"
	"dbconnection/test/order"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

func PreviousOrder(db *gorm.DB, id string) {
	fmt.Printf("%-20s %-20s %-20s %-20s\n", "Item name", "|Item Desc", "|Quantity", "|Cost/unit")
	var orders []*order.Order
	db.Where("customer_id = ? && is_paid = ?", id, true).Find(&orders)
	for _, order := range orders {
		fmt.Printf("%-20s %-20s %-20d %-20f\n", order.ItemName, order.ItemDesc, order.Quantity, order.CostPerUnit)
	}
	if len(orders) == 0 {
		fmt.Println("You haven't order any items")
	}
	var tempNum string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter 0 for Dashboard page")
		if scanner.Scan() {
			tempNum = scanner.Text()
		}
		flag, err := strconv.ParseBool(tempNum)
		if err != nil {
			fmt.Println(errors.New("You have inserted invalid value"))
		} else if flag != true {
			break
		} else {
			fmt.Println("Enter valid number, For Back to dashboard Enter 0")
		}
	}
}
