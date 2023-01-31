package requests

type SAPSalesOrderHeaderPartner struct {
	SalesOrder      string  `json:"SalesOrder"`
	PartnerFunction string  `json:"PartnerFunction"`
	Customer        *string `json:"Customer"`
	Supplier        *string `json:"Supplier"`
}
