package requests

type ConversionData struct {
	SalesOrder     string  `json:"SalesOrder"`
	OrderID        int     `json:"OrderID"`
	SalesOrderItem string  `json:"SalesOrderItem"`
	OrderItem      int     `json:"OrderItem"`
	Buyer          int     `json:"Buyer"`
	SoldToParty    *string `json:"SoldToParty"`
	Material       string  `json:"Material"`
	Product        *string `json:"Product"`
}
