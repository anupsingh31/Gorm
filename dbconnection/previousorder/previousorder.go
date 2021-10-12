package previousorder

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func PreviousOrder(db *gorm.DB, id uint) {
	fmt.Println("Item name\t\t| Item Desc\t\t| Quantity\t\t| Cost/unit\t\t| ")
	rows, err := db.DB().Query("SELECT item_name,item_desc,quantity,cost_per_unit FROM orders WHERE customer_id = ? && is_paid = ?", id, true)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var itemName, itemDesc string
		var quantity int
		var cost float64
		err = rows.Scan(&itemName, &itemDesc, &quantity, &cost)
		if err != nil {
			panic(err)
		}
		fmt.Println(itemName, "\t| ", itemDesc, "\t|", quantity, "\t|", cost, "\t|")
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
