package requests

type CodeConversionItem struct {
	SalesOrderItem    string  `json:"SalesOrderItem"`
	OrderItem         int     `json:"OrderItem"`
	OrderItemCategory *string `json:"OrderItemCategory"`
	Material          string  `json:"Material"`
	Product           *string `json:"Product"`
	ProductGroup      *string `json:"ProductGroup"`
	DeliverToParty    *int    `json:"DeliverToParty"`
}
