package order

import (
	"dbconnection/test/model"
)

//var Id uint = 1

type Order struct {
	model.Model
	CustomerID  string `sql:"REFERENCES customers(id)" type:"uuid"`
	ItemName    string
	ItemDesc    string
	Quantity    int
	CostPerUnit float64
	IsPaid      *bool
}

func New(ID string, oItemName, oItemDesc string, oQuantity int, oCostPerUnit float64, oIspaid *bool) *Order {
	return &Order{
		CustomerID:  ID,
		ItemName:    oItemName,
		ItemDesc:    oItemDesc,
		Quantity:    oQuantity,
		CostPerUnit: oCostPerUnit,
		IsPaid:      oIspaid,
	}
}

// func (o *Order) CreateOrderTable() {
// 	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/swabhav?charset=utf8&parseTime=True&loc=Local")
// 	defer db.Close()
// 	if err != nil {
// 		log.Println("Connection Failed to Open")
// 	} else {
// 		log.Println("Connection Established")
// 	}
// 	db.AutoMigrate(&Order{})
// 	db.Create(o)
// }
