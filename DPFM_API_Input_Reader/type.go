package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey       string              `json:"connection_key"`
	Result              bool                `json:"result"`
	RedisKey            string              `json:"redis_key"`
	Filepath            string              `json:"filepath"`
	APIStatusCode       int                 `json:"api_status_code"`
	RuntimeSessionID    string              `json:"runtime_session_id"`
	BusinessPartnerID   *int                `json:"business_partner"`
	ServiceLabel        string              `json:"service_label"`
	SAPSalesOrderHeader SAPSalesOrderHeader `json:"SalesOrder"`
	APISchema           string              `json:"api_schema"`
	Accepter            []string            `json:"accepter"`
	Deleted             bool                `json:"deleted"`
}

type SAPSalesOrderHeader struct {
	SalesOrder                     string                       `json:"SalesOrder"`
	SalesOrderType                 *string                      `json:"SalesOrderType"`
	SalesOrganization              *string                      `json:"SalesOrganization"`
	DistributionChannel            *string                      `json:"DistributionChannel"`
	OrganizationDivision           *string                      `json:"OrganizationDivision"`
	SalesGroup                     *string                      `json:"SalesGroup"`
	SalesOffice                    *string                      `json:"SalesOffice"`
	SalesDistrict                  *string                      `json:"SalesDistrict"`
	SoldToParty                    *string                      `json:"SoldToParty"`
	CreationDate                   *string                      `json:"CreationDate"`
	LastChangeDate                 *string                      `json:"LastChangeDate"`
	ExternalDocumentID             *string                      `json:"ExternalDocumentID"`
	LastChangeDateTime             *string                      `json:"LastChangeDateTime"`
	PurchaseOrderByCustomer        *string                      `json:"PurchaseOrderByCustomer"`
	CustomerPurchaseOrderDate      *string                      `json:"CustomerPurchaseOrderDate"`
	SalesOrderDate                 *string                      `json:"SalesOrderDate"`
	TotalNetAmount                 *string                      `json:"TotalNetAmount"`
	OverallDeliveryStatus          *string                      `json:"OverallDeliveryStatus"`
	TotalBlockStatus               *string                      `json:"TotalBlockStatus"`
	OverallOrdReltdBillgStatus     *string                      `json:"OverallOrdReltdBillgStatus"`
	OverallSDDocReferenceStatus    *string                      `json:"OverallSDDocReferenceStatus"`
	TransactionCurrency            *string                      `json:"TransactionCurrency"`
	SDDocumentReason               *string                      `json:"SDDocumentReason"`
	PricingDate                    *string                      `json:"PricingDate"`
	PriceDetnExchangeRate          *string                      `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate          *string                      `json:"RequestedDeliveryDate"`
	ShippingCondition              *string                      `json:"ShippingCondition"`
	CompleteDeliveryIsDefined      *bool                        `json:"CompleteDeliveryIsDefined"`
	ShippingType                   *string                      `json:"ShippingType"`
	HeaderBillingBlockReason       *string                      `json:"HeaderBillingBlockReason"`
	DeliveryBlockReason            *string                      `json:"DeliveryBlockReason"`
	IncotermsClassification        *string                      `json:"IncotermsClassification"`
	CustomerPriceGroup             *string                      `json:"CustomerPriceGroup"`
	PriceListType                  *string                      `json:"PriceListType"`
	CustomerPaymentTerms           *string                      `json:"CustomerPaymentTerms"`
	PaymentMethod                  *string                      `json:"PaymentMethod"`
	ReferenceSDDocument            *string                      `json:"ReferenceSDDocument"`
	ReferenceSDDocumentCategory    *string                      `json:"ReferenceSDDocumentCategory"`
	CustomerAccountAssignmentGroup *string                      `json:"CustomerAccountAssignmentGroup"`
	AccountingExchangeRate         *string                      `json:"AccountingExchangeRate"`
	CustomerGroup                  *string                      `json:"CustomerGroup"`
	AdditionalCustomerGroup1       *string                      `json:"AdditionalCustomerGroup1"`
	AdditionalCustomerGroup2       *string                      `json:"AdditionalCustomerGroup2"`
	AdditionalCustomerGroup3       *string                      `json:"AdditionalCustomerGroup3"`
	AdditionalCustomerGroup4       *string                      `json:"AdditionalCustomerGroup4"`
	AdditionalCustomerGroup5       *string                      `json:"AdditionalCustomerGroup5"`
	CustomerTaxClassification1     *string                      `json:"CustomerTaxClassification1"`
	TotalCreditCheckStatus         *string                      `json:"TotalCreditCheckStatus"`
	BillingDocumentDate            *string                      `json:"BillingDocumentDate"`
	SAPSalesOrderHeaderPartner     []SAPSalesOrderHeaderPartner `json:"HeaderPartner"`
	SAPSalesOrderItem              []SAPSalesOrderItem          `json:"SalesOrderItem"`
}

type SAPSalesOrderHeaderPartner struct {
	SalesOrder      string  `json:"SalesOrder"`
	PartnerFunction string  `json:"PartnerFunction"`
	Customer        *string `json:"Customer"`
	Supplier        *string `json:"Supplier"`
}

type SAPSalesOrderItem struct {
	SalesOrder                      string                            `json:"SalesOrder"`
	SalesOrderItem                  string                            `json:"SalesOrderItem"`
	SalesOrderItemCategory          *string                           `json:"SalesOrderItemCategory"`
	SalesOrderItemText              *string                           `json:"SalesOrderItemText"`
	PurchaseOrderByCustomer         *string                           `json:"PurchaseOrderByCustomer"`
	Material                        *string                           `json:"Material"`
	MaterialByCustomer              *string                           `json:"MaterialByCustomer"`
	PricingDate                     *string                           `json:"PricingDate"`
	BillingPlan                     *string                           `json:"BillingPlan"`
	RequestedQuantity               *string                           `json:"RequestedQuantity"`
	RequestedQuantityUnit           *string                           `json:"RequestedQuantityUnit"`
	OrderQuantityUnit               *string                           `json:"OrderQuantityUnit"`
	ConfdDelivQtyInOrderQtyUnit     *string                           `json:"ConfdDelivQtyInOrderQtyUnit"`
	ItemGrossWeight                 *string                           `json:"ItemGrossWeight"`
	ItemNetWeight                   *string                           `json:"ItemNetWeight"`
	ItemWeightUnit                  *string                           `json:"ItemWeightUnit"`
	ItemVolume                      *string                           `json:"ItemVolume"`
	ItemVolumeUnit                  *string                           `json:"ItemVolumeUnit"`
	TransactionCurrency             *string                           `json:"TransactionCurrency"`
	NetAmount                       *string                           `json:"NetAmount"`
	MaterialGroup                   *string                           `json:"MaterialGroup"`
	MaterialPricingGroup            *string                           `json:"MaterialPricingGroup"`
	BillingDocumentDate             *string                           `json:"BillingDocumentDate"`
	Batch                           *string                           `json:"Batch"`
	ProductionPlant                 *string                           `json:"ProductionPlant"`
	StorageLocation                 *string                           `json:"StorageLocation"`
	DeliveryGroup                   *string                           `json:"DeliveryGroup"`
	ShippingPoint                   *string                           `json:"ShippingPoint"`
	ShippingType                    *string                           `json:"ShippingType"`
	DeliveryPriority                *string                           `json:"DeliveryPriority"`
	IncotermsClassification         *string                           `json:"IncotermsClassification"`
	TaxAmount                       *string                           `json:"TaxAmount"`
	TaxCode                         *string                           `json:"TaxCode"`
	ProductTaxClassification1       *string                           `json:"ProductTaxClassification1"`
	MatlAccountAssignmentGroup      *string                           `json:"MatlAccountAssignmentGroup"`
	GrossAmount                     *string                           `json:"GrossAmount"`
	CostAmount                      *string                           `json:"CostAmount"`
	CustomerPaymentTerms            *string                           `json:"CustomerPaymentTerms"`
	CustomerGroup                   *string                           `json:"CustomerGroup"`
	SalesDocumentRjcnReason         *string                           `json:"SalesDocumentRjcnReason"`
	ItemBillingBlockReason          *string                           `json:"ItemBillingBlockReason"`
	WBSElement                      *string                           `json:"WBSElement"`
	ProfitCenter                    *string                           `json:"ProfitCenter"`
	AccountingExchangeRate          *string                           `json:"AccountingExchangeRate"`
	ReferenceSDDocument             *string                           `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem         *string                           `json:"ReferenceSDDocumentItem"`
	SDProcessStatus                 *string                           `json:"SDProcessStatus"`
	DeliveryStatus                  *string                           `json:"DeliveryStatus"`
	OrderRelatedBillingStatus       *string                           `json:"OrderRelatedBillingStatus"`
	SAPSalesOrderItemPartner        SAPSalesOrderItemPartner          `json:"ItemPartner"`
	SAPSalesOrderItemPricingElement []SAPSalesOrderItemPricingElement `json:"ItemPricingElement"`
	SAPSalesOrderItemScheduleLine   []SAPSalesOrderItemScheduleLine   `json:"ItemScheduleLine"`
}

type SAPSalesOrderItemPartner struct {
	SalesOrder      string  `json:"SalesOrder"`
	SalesOrderItem  string  `json:"SalesOrderItem"`
	PartnerFunction string  `json:"PartnerFunction"`
	Customer        *string `json:"Customer"`
	Supplier        *string `json:"Supplier"`
	Personnel       *string `json:"Personnel"`
	ContactPerson   *string `json:"ContactPerson"`
}

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

type SAPSalesOrderItemScheduleLine struct {
	SalesOrder                    string  `json:"SalesOrder"`
	SalesOrderItem                string  `json:"SalesOrderItem"`
	ScheduleLine                  string  `json:"ScheduleLine"`
	RequestedDeliveryDate         *string `json:"RequestedDeliveryDate"`
	ConfirmedDeliveryDate         *string `json:"ConfirmedDeliveryDate"`
	OrderQuantityUnit             *string `json:"OrderQuantityUnit"`
	ScheduleLineOrderQuantity     *string `json:"ScheduleLineOrderQuantity"`
	ConfdOrderQtyByMatlAvailCheck *string `json:"ConfdOrderQtyByMatlAvailCheck"`
	DeliveredQtyInOrderQtyUnit    *string `json:"DeliveredQtyInOrderQtyUnit"`
	OpenConfdDelivQtyInOrdQtyUnit *string `json:"OpenConfdDelivQtyInOrdQtyUnit"`
	CorrectedQtyInOrderQtyUnit    *string `json:"CorrectedQtyInOrderQtyUnit"`
	DelivBlockReasonForSchedLine  *string `json:"DelivBlockReasonForSchedLine"`
}
