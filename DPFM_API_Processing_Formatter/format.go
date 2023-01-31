package dpfm_api_processing_formatter

import (
	"convert-to-dpfm-orders-from-sap-sales-order/DPFM_API_Caller/requests"
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-sales-order/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
	"strconv"
)

// データ連携基盤とSAP Sales Orderとの項目マッピング
// Header
func (psdc *SDC) ConvertToMappingHeader(sdc *dpfm_api_input_reader.SDC) (*MappingHeader, error) {
	data := sdc.SAPSalesOrderHeader

	systemDate := GetSystemDatePtr()

	priceDetnExchangeRate, err := ParseFloat32Ptr(data.PriceDetnExchangeRate)
	if err != nil {
		return nil, err
	}
	accountingExchangeRate, err := ParseFloat32Ptr(data.AccountingExchangeRate)
	if err != nil {
		return nil, err
	}
	headerBlockStatus, err := ParseBoolPtr(data.TotalBlockStatus)
	if err != nil {
		return nil, err
	}
	headerBillingBlockStatus, err := ParseBoolPtr(data.HeaderBillingBlockReason)
	if err != nil {
		return nil, err
	}
	headerDeliveryBlockStatus, err := ParseBoolPtr(data.DeliveryBlockReason)
	if err != nil {
		return nil, err
	}

	res := MappingHeader{
		OrderDate:                 data.SalesOrderDate,
		Seller:                    sdc.BusinessPartnerID,
		BillFromParty:             sdc.BusinessPartnerID,
		Payee:                     sdc.BusinessPartnerID,
		CreationDate:              systemDate,
		LastChangeDate:            systemDate,
		TransactionCurrency:       data.TransactionCurrency,
		PricingDate:               data.PricingDate,
		PriceDetnExchangeRate:     priceDetnExchangeRate,
		RequestedDeliveryDate:     data.RequestedDeliveryDate,
		Incoterms:                 data.IncotermsClassification,
		PaymentTerms:              data.CustomerPaymentTerms,
		PaymentMethod:             data.PaymentMethod,
		AccountingExchangeRate:    accountingExchangeRate,
		InvoiceDocumentDate:       data.BillingDocumentDate,
		HeaderBlockStatus:         headerBlockStatus,
		HeaderBillingBlockStatus:  headerBillingBlockStatus,
		HeaderDeliveryBlockStatus: headerDeliveryBlockStatus,
	}

	return &res, nil
}

// Item
func (psdc *SDC) ConvertToMappingItem(sdc *dpfm_api_input_reader.SDC) (*[]MappingItem, error) {
	var res []MappingItem
	data := sdc.SAPSalesOrderHeader
	dataItem := sdc.SAPSalesOrderHeader.SAPSalesOrderItem

	systemDate := GetSystemDatePtr()

	for _, dataItem := range dataItem {
		priceDetnExchangeRate, err := ParseFloat32Ptr(data.PriceDetnExchangeRate)
		if err != nil {
			return nil, err
		}
		requestedQuantity, err := ParseFloat32Ptr(dataItem.RequestedQuantity)
		if err != nil {
			return nil, err
		}
		confdDelivQtyInOrderQtyUnit, err := ParseFloat32Ptr(dataItem.ConfdDelivQtyInOrderQtyUnit)
		if err != nil {
			return nil, err
		}
		itemWeightUnit, err := ParseFloat32Ptr(dataItem.ItemWeightUnit)
		if err != nil {
			return nil, err
		}
		itemGrossWeight, err := ParseFloat32Ptr(dataItem.ItemGrossWeight)
		if err != nil {
			return nil, err
		}
		itemNetWeight, err := ParseFloat32Ptr(dataItem.ItemNetWeight)
		if err != nil {
			return nil, err
		}
		netAmount, err := ParseFloat32Ptr(dataItem.NetAmount)
		if err != nil {
			return nil, err
		}
		taxAmount, err := ParseFloat32Ptr(dataItem.TaxAmount)
		if err != nil {
			return nil, err
		}
		grossAmount, err := ParseFloat32Ptr(dataItem.GrossAmount)
		if err != nil {
			return nil, err
		}

		res = append(res, MappingItem{
			SalesOrder:                              data.SalesOrder,
			OrderItemTextBySeller:                   dataItem.SalesOrderItemText,
			BaseUnit:                                dataItem.RequestedQuantityUnit,
			PricingDate:                             dataItem.PricingDate,
			PriceDetnExchangeRate:                   priceDetnExchangeRate,
			RequestedDeliveryDate:                   data.RequestedDeliveryDate,
			DeliverFromParty:                        sdc.BusinessPartnerID,
			CreationDate:                            systemDate,
			LastChangeDate:                          systemDate,
			DeliverFromPlant:                        dataItem.ShippingPoint,
			DeliverFromPlantStorageLocation:         dataItem.StorageLocation,
			DeliveryUnit:                            dataItem.OrderQuantityUnit,
			StockConfirmationBusinessPartner:        sdc.BusinessPartnerID,
			StockConfirmationPlant:                  dataItem.ProductionPlant,
			StockConfirmationPlantBatch:             dataItem.Batch,
			OrderQuantityInBaseUnit:                 requestedQuantity,
			ConfirmedOrderQuantityInBaseUnit:        confdDelivQtyInOrderQtyUnit,
			ItemWeightUnit:                          itemWeightUnit,
			ItemGrossWeight:                         itemGrossWeight,
			ItemNetWeight:                           itemNetWeight,
			NetAmount:                               netAmount,
			TaxAmount:                               taxAmount,
			GrossAmount:                             grossAmount,
			InvoiceDocumentDate:                     dataItem.BillingDocumentDate,
			ProductionPlantBusinessPartner:          sdc.BusinessPartnerID,
			ProductionPlant:                         dataItem.ProductionPlant,
			Incoterms:                               dataItem.IncotermsClassification,
			TransactionTaxClassification:            data.CustomerTaxClassification1,
			ProductTaxClassificationBillToCountry:   dataItem.ProductTaxClassification1,
			ProductTaxClassificationBillFromCountry: dataItem.ProductTaxClassification1,
			AccountAssignmentGroup:                  data.CustomerAccountAssignmentGroup,
			ProductAccountAssignmentGroup:           dataItem.MatlAccountAssignmentGroup,
			PaymentTerms:                            dataItem.CustomerPaymentTerms,
			TaxCode:                                 dataItem.TaxCode,
		})
	}

	return &res, nil
}

// ItemPricingElement
func (psdc *SDC) ConvertToMappingItemPricingElement(sdc *dpfm_api_input_reader.SDC) (*[]MappingItemPricingElement, error) {
	var res []MappingItemPricingElement
	data := sdc.SAPSalesOrderHeader
	dataItem := sdc.SAPSalesOrderHeader.SAPSalesOrderItem

	for _, dataItem := range dataItem {
		dataItemPricingElement := dataItem.SAPSalesOrderItemPricingElement
		for _, dataItemPricingElement := range dataItemPricingElement {
			conditionRateValue, err := ParseFloat32Ptr(dataItemPricingElement.ConditionRateValue)
			if err != nil {
				return nil, err
			}
			conditionQuantity, err := ParseFloat32Ptr(dataItemPricingElement.ConditionQuantity)
			if err != nil {
				return nil, err
			}
			conditionAmount, err := ParseFloat32Ptr(dataItemPricingElement.ConditionAmount)
			if err != nil {
				return nil, err
			}

			res = append(res, MappingItemPricingElement{
				SalesOrder:            data.SalesOrder,
				SalesOrderItem:        dataItem.SalesOrderItem,
				SoldToParty:           data.SoldToParty,
				PricingDate:           dataItem.PricingDate,
				ConditionRateValue:    conditionRateValue,
				ConditionCurrency:     dataItemPricingElement.ConditionCurrency,
				ConditionQuantity:     conditionQuantity,
				ConditionQuantityUnit: dataItemPricingElement.ConditionQuantityUnit,
				TaxCode:               dataItemPricingElement.TaxCode,
				ConditionAmount:       conditionAmount,
				TransactionCurrency:   dataItemPricingElement.TransactionCurrency,
			})
		}
	}

	return &res, nil
}

func (psdc *SDC) ConvertToMappingItemScheduleLine(sdc *dpfm_api_input_reader.SDC) (*[]MappingItemScheduleLine, error) {
	var res []MappingItemScheduleLine
	data := sdc.SAPSalesOrderHeader
	dataItem := sdc.SAPSalesOrderHeader.SAPSalesOrderItem

	for _, dataItem := range dataItem {
		dataItemScheduleLine := dataItem.SAPSalesOrderItemScheduleLine
		for _, dataItemScheduleLine := range dataItemScheduleLine {
			scheduleLineOrderQuantity, err := ParseFloat32Ptr(dataItemScheduleLine.ScheduleLineOrderQuantity)
			if err != nil {
				return nil, err
			}
			confdOrderQtyByMatlAvailCheck, err := ParseFloat32Ptr(dataItemScheduleLine.ConfdOrderQtyByMatlAvailCheck)
			if err != nil {
				return nil, err
			}
			deliveredQtyInOrderQtyUnit, err := ParseFloat32Ptr(dataItemScheduleLine.DeliveredQtyInOrderQtyUnit)
			if err != nil {
				return nil, err
			}
			openConfdDelivQtyInOrdQtyUnit, err := ParseFloat32Ptr(dataItemScheduleLine.OpenConfdDelivQtyInOrdQtyUnit)
			if err != nil {
				return nil, err
			}

			res = append(res, MappingItemScheduleLine{
				SalesOrder:                            data.SalesOrder,
				SalesOrderItem:                        dataItem.SalesOrderItem,
				Material:                              *dataItem.Material,
				StockConfirmationBussinessPartner:     sdc.BusinessPartnerID,
				StockConfirmationPlant:                dataItem.ProductionPlant,
				StockConfirmationPlantBatch:           dataItem.Batch,
				RequestedDeliveryDate:                 dataItemScheduleLine.RequestedDeliveryDate,
				ConfirmedDeliveryDate:                 dataItemScheduleLine.ConfirmedDeliveryDate,
				OrderQuantityInBaseUnit:               scheduleLineOrderQuantity,
				ConfirmedOrderQuantityByPDTAvailCheck: confdOrderQtyByMatlAvailCheck,
				DeliveredQuantityInBaseUnit:           deliveredQtyInOrderQtyUnit,
				OpenConfirmedQuantityInBaseUnit:       openConfdDelivQtyInOrdQtyUnit,
			})
		}
	}

	return &res, nil
}

// Address
func (psdc *SDC) ConvertToMappingAddress(sdc *dpfm_api_input_reader.SDC) (*MappingAddress, error) {
	data := sdc.SAPSalesOrderHeader

	res := MappingAddress{
		SalesOrder: data.SalesOrder,
	}

	return &res, nil
}

// Partner
func (psdc *SDC) ConvertToMappingPartner(sdc *dpfm_api_input_reader.SDC) (*[]MappingPartner, error) {
	var res []MappingPartner
	data := sdc.SAPSalesOrderHeader
	dataPartner := sdc.SAPSalesOrderHeader.SAPSalesOrderHeaderPartner

	for range dataPartner {
		res = append(res, MappingPartner{
			SalesOrder:         data.SalesOrder,
			ExternalDocumentID: data.PurchaseOrderByCustomer,
		})
	}

	return &res, nil
}

// 番号処理
func (psdc *SDC) ConvertToCodeConversionKey(sdc *dpfm_api_input_reader.SDC, labelConvertFrom, labelConvertTo string, codeConvertFrom any) *CodeConversionKey {
	pm := &requests.CodeConversionKey{
		SystemConvertTo:   "DPFM",
		SystemConvertFrom: "SAP",
		BusinessPartner:   *sdc.BusinessPartnerID,
	}

	pm.LabelConvertFrom = labelConvertFrom
	pm.LabelConvertTo = labelConvertTo

	switch codeConvertFrom := codeConvertFrom.(type) {
	case string:
		pm.CodeConvertFrom = codeConvertFrom
	case int:
		pm.CodeConvertFrom = strconv.FormatInt(int64(codeConvertFrom), 10)
	case float32:
		pm.CodeConvertFrom = strconv.FormatFloat(float64(codeConvertFrom), 'f', -1, 32)
	case bool:
		pm.CodeConvertFrom = strconv.FormatBool(codeConvertFrom)
	case *string:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = *codeConvertFrom
		}
	case *int:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatInt(int64(*codeConvertFrom), 10)
		}
	case *float32:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatFloat(float64(*codeConvertFrom), 'f', -1, 32)
		}
	case *bool:
		if codeConvertFrom != nil {
			pm.CodeConvertFrom = strconv.FormatBool(*codeConvertFrom)
		}
	}

	data := pm
	res := CodeConversionKey{
		SystemConvertTo:   data.SystemConvertTo,
		SystemConvertFrom: data.SystemConvertFrom,
		LabelConvertTo:    data.LabelConvertTo,
		LabelConvertFrom:  data.LabelConvertFrom,
		CodeConvertFrom:   data.CodeConvertFrom,
		BusinessPartner:   data.BusinessPartner,
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionQueryGets(rows *sql.Rows) (*[]CodeConversionQueryGets, error) {
	defer rows.Close()
	var res []CodeConversionQueryGets

	i := 0
	for rows.Next() {
		i++
		pm := &requests.CodeConversionQueryGets{}

		err := rows.Scan(
			&pm.CodeConversionID,
			&pm.SystemConvertTo,
			&pm.SystemConvertFrom,
			&pm.LabelConvertTo,
			&pm.LabelConvertFrom,
			&pm.CodeConvertFrom,
			&pm.CodeConvertTo,
			&pm.BusinessPartner,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, CodeConversionQueryGets{
			CodeConversionID:  data.CodeConversionID,
			SystemConvertTo:   data.SystemConvertTo,
			SystemConvertFrom: data.SystemConvertFrom,
			LabelConvertTo:    data.LabelConvertTo,
			LabelConvertFrom:  data.LabelConvertFrom,
			CodeConvertFrom:   data.CodeConvertFrom,
			CodeConvertTo:     data.CodeConvertTo,
			BusinessPartner:   data.BusinessPartner,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_code_conversion_code_conversion_data'テーブルに対象のレコードが存在しません。")
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionHeader(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionHeader, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionHeader{}

	pm.SalesOrder = dataMap["OrderID"].CodeConvertFrom
	pm.OrderID, err = ParseInt(dataMap["OrderID"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.OrderType = GetStringPtr(dataMap["OrderType"].CodeConvertTo)
	pm.SoldToParty = GetStringPtr(dataMap["Buyer"].CodeConvertFrom)
	pm.Buyer, err = ParseIntPtr(GetStringPtr(dataMap["Buyer"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.BillToParty, err = ParseIntPtr(GetStringPtr(dataMap["BillToParty"].CodeConvertTo))
	if err != nil {
		return nil, err
	}
	pm.Payer, err = ParseIntPtr(GetStringPtr(dataMap["Payer"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionHeader{
		SalesOrder:  data.SalesOrder,
		OrderID:     data.OrderID,
		OrderType:   data.OrderType,
		SoldToParty: data.SoldToParty,
		Buyer:       data.Buyer,
		BillToParty: data.BillToParty,
		Payer:       data.Payer,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionPartner(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionPartner, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionPartner{}

	pm.PartnerFunction = dataMap["PartnerFunction"].CodeConvertTo
	pm.BusinessPartner, err = ParseInt(dataMap["BusinessPartner"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionPartner{
		PartnerFunction: data.PartnerFunction,
		BusinessPartner: data.BusinessPartner,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionItem(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItem, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItem{}

	pm.SalesOrderItem = dataMap["OrderItem"].CodeConvertFrom
	pm.OrderItem, err = ParseInt(dataMap["OrderItem"].CodeConvertTo)
	if err != nil {
		return nil, err
	}
	pm.OrderItemCategory = GetStringPtr(dataMap["OrderItemCategory"].CodeConvertTo)
	pm.Material = dataMap["Product"].CodeConvertFrom
	pm.Product = GetStringPtr(dataMap["Product"].CodeConvertTo)
	pm.ProductGroup = GetStringPtr(dataMap["ProductGroup"].CodeConvertTo)
	pm.DeliverToParty, err = ParseIntPtr(GetStringPtr(dataMap["DeliverToParty"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItem{
		SalesOrderItem:    data.SalesOrderItem,
		OrderItem:         data.OrderItem,
		OrderItemCategory: data.OrderItemCategory,
		Material:          data.Material,
		Product:           data.Product,
		ProductGroup:      data.ProductGroup,
		DeliverToParty:    data.DeliverToParty,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToConversionData() *[]ConversionData {
	var res []ConversionData

	for _, v := range *psdc.CodeConversionItem {
		pm := &requests.ConversionData{}

		pm.SalesOrder = psdc.CodeConversionHeader.SalesOrder
		pm.OrderID = psdc.CodeConversionHeader.OrderID
		pm.SalesOrderItem = v.SalesOrderItem
		pm.OrderItem = v.OrderItem
		pm.SoldToParty = psdc.CodeConversionHeader.SoldToParty
		pm.Buyer = *psdc.CodeConversionHeader.Buyer
		pm.Material = v.Material
		pm.Product = v.Product

		data := pm
		res = append(res, ConversionData{
			SalesOrder:     data.SalesOrder,
			OrderID:        data.OrderID,
			SalesOrderItem: data.SalesOrderItem,
			OrderItem:      data.OrderItem,
			SoldToParty:    data.SoldToParty,
			Buyer:          data.Buyer,
			Material:       data.Material,
			Product:        data.Product,
		})
	}

	return &res
}

func (psdc *SDC) ConvertToCodeConversionItemPricingElement(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItemPricingElement, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItemPricingElement{}

	pm.ConditionSequentialNumber, err = ParseIntPtr(GetStringPtr(dataMap["ConditionSequentialNumber"].CodeConvertTo))
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItemPricingElement{
		ConditionSequentialNumber: data.ConditionSequentialNumber,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCodeConversionItemScheduleLine(dataQueryGets *[]CodeConversionQueryGets) (*CodeConversionItemScheduleLine, error) {
	var err error

	dataMap := make(map[string]CodeConversionQueryGets, len(*dataQueryGets))
	for _, v := range *dataQueryGets {
		dataMap[v.LabelConvertTo] = v
	}

	pm := &requests.CodeConversionItemScheduleLine{}

	pm.ScheduleLine, err = ParseInt(dataMap["ScheduleLine"].CodeConvertTo)
	if err != nil {
		return nil, err
	}

	data := pm
	res := CodeConversionItemScheduleLine{
		ScheduleLine: data.ScheduleLine,
	}

	return &res, nil
}

// 個別処理
func (psdc *SDC) ConvertToSetCustomer(headerPartner dpfm_api_input_reader.SAPSalesOrderHeaderPartner) *SetCustomer {
	pm := &requests.SetCustomer{}

	pm.PartnerFunction = headerPartner.PartnerFunction
	pm.Customer = headerPartner.Customer

	data := pm
	res := SetCustomer{
		PartnerFunction: data.PartnerFunction,
		Customer:        data.Customer,
	}

	return &res
}

func (psdc *SDC) ConvertToSetCustomerHeader(customer []SetCustomer) *SetCustomerHeader {
	pm := &requests.SetCustomerHeader{}

	customerMap := make(map[string]SetCustomer, len(customer))
	for _, v := range customer {
		customerMap[v.PartnerFunction] = v
	}

	pm.CustomerForBillToParty = customerMap["BL"].Customer
	pm.CustomerForPayer = customerMap["PY"].Customer

	data := pm
	res := SetCustomerHeader{
		CustomerForBillToParty: data.CustomerForBillToParty,
		CustomerForPayer:       data.CustomerForPayer,
	}

	return &res
}

func (psdc *SDC) ConvertToSetCustomerItem(customer []SetCustomer) *SetCustomerItem {
	pm := &requests.SetCustomerItem{}

	customerMap := make(map[string]SetCustomer, len(customer))
	for _, v := range customer {
		customerMap[v.PartnerFunction] = v
	}

	pm.CustomerForDeliverToParty = customerMap["SH"].Customer

	data := pm
	res := SetCustomerItem{
		CustomerForDeliverToParty: data.CustomerForDeliverToParty,
	}

	return &res
}

func (psdc *SDC) ConvertToSetPartnerFunction(partnerFunction string) *SetPartnerFunction {
	pm := &requests.SetPartnerFunction{}

	pm.PartnerFunction = partnerFunction

	data := pm
	res := SetPartnerFunction{
		PartnerFunction: data.PartnerFunction,
	}

	return &res
}

func (psdc *SDC) ConvertToSetConditionType(conditionType *string) *SetConditionType {
	pm := &requests.SetConditionType{}

	pm.ConditionType = conditionType

	data := pm
	res := SetConditionType{
		ConditionType: data.ConditionType,
	}

	return &res
}
