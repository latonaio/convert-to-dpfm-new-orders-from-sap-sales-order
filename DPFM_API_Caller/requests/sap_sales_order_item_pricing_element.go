package requests

type SAPSalesOrderItemPricingElement struct {
	SalesOrder                     string  `json:"SalesOrder"`
	SalesOrderItem                 string  `json:"SalesOrderItem"`
	PricingProcedureStep           string  `json:"PricingProcedureStep"`
	PricingProcedureCounter        string  `json:"PricingProcedureCounter"`
	ConditionType                  *string `json:"ConditionType"`
	PriceConditionDeterminationDte *string `json:"PriceConditionDeterminationDte"`
	ConditionCalculationType       *string `json:"ConditionCalculationType"`
	ConditionBaseValue             *string `json:"ConditionBaseValue"`
	ConditionRateValue             *string `json:"ConditionRateValue"`
	ConditionCurrency              *string `json:"ConditionCurrency"`
	ConditionQuantity              *string `json:"ConditionQuantity"`
	ConditionQuantityUnit          *string `json:"ConditionQuantityUnit"`
	ConditionCategory              *string `json:"ConditionCategory"`
	PricingScaleType               *string `json:"PricingScaleType"`
	ConditionRecord                *string `json:"ConditionRecord"`
	ConditionSequentialNumber      *string `json:"ConditionSequentialNumber"`
	TaxCode                        *string `json:"TaxCode"`
	ConditionAmount                *string `json:"ConditionAmount"`
	TransactionCurrency            *string `json:"TransactionCurrency"`
	PricingScaleBasis              *string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue       *string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit        *string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency    *string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged     *bool   `json:"ConditionIsManuallyChanged"`
}
