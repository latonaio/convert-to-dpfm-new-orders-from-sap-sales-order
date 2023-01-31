package requests

type CodeConversionHeader struct {
	SalesOrder  string  `json:"SalesOrder"`
	OrderID     int     `json:"OrderID"`
	OrderType   *string `json:"OrderType"`
	SoldToParty *string `json:"SoldToParty"`
	Buyer       *int    `json:"Buyer"`
	BillToParty *int    `json:"BillToParty"`
	Payer       *int    `json:"Payer"`
}
