package convert_complementer

import (
	dpfm_api_input_reader "convert-to-dpfm-orders-from-sap-sales-order/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "convert-to-dpfm-orders-from-sap-sales-order/DPFM_API_Processing_Formatter"
	"strings"
)

// Mapping Itemの処理
func (c *ConvertComplementer) ComplementMappingItem(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.MappingItem, error) {
	res, err := psdc.ConvertToMappingItem(sdc)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ConvertComplementer) CodeConversionItem(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) (*[]dpfm_api_processing_formatter.CodeConversionItem, error) {
	var data []dpfm_api_processing_formatter.CodeConversionItem

	for _, item := range sdc.SAPSalesOrderHeader.SAPSalesOrderItem {
		var dataKey []*dpfm_api_processing_formatter.CodeConversionKey
		var args []interface{}

		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "SalesOrderItem", "OrderItem", item.SalesOrderItem))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "SalesOrderItemCategory", "OrderItemCategory", item.SalesOrderItemCategory))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Material", "Product", item.Material))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "MaterialGroup", "ProductGroup", item.MaterialGroup))
		dataKey = append(dataKey, psdc.ConvertToCodeConversionKey(sdc, "Customer", "DeliverToParty", psdc.SetCustomerItem.CustomerForDeliverToParty))

		repeat := strings.Repeat("(?,?,?,?,?,?),", len(dataKey)-1) + "(?,?,?,?,?,?)"
		for _, v := range dataKey {
			args = append(args, v.SystemConvertTo, v.SystemConvertFrom, v.LabelConvertTo, v.LabelConvertFrom, v.CodeConvertFrom, v.BusinessPartner)
		}

		rows, err := c.db.Query(
			`SELECT CodeConversionID, SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom,
			CodeConvertFrom, CodeConvertTo, BusinessPartner
			FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_code_conversion_code_conversion_data
			WHERE (SystemConvertTo, SystemConvertFrom, LabelConvertTo, LabelConvertFrom, CodeConvertFrom, BusinessPartner) IN ( `+repeat+` );`, args...,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		dataQueryGets, err := psdc.ConvertToCodeConversionQueryGets(rows)
		if err != nil {
			return nil, err
		}

		datum, err := psdc.ConvertToCodeConversionItem(dataQueryGets)
		if err != nil {
			return nil, err
		}

		data = append(data, *datum)
	}

	return &data, nil
}

func (c *ConvertComplementer) SetCustomerItem(sdc *dpfm_api_input_reader.SDC, psdc *dpfm_api_processing_formatter.SDC) *dpfm_api_processing_formatter.SetCustomerItem {
	var dataCustomer []dpfm_api_processing_formatter.SetCustomer

	for _, v := range sdc.SAPSalesOrderHeader.SAPSalesOrderHeaderPartner {
		partnerFunction := v.PartnerFunction
		if partnerFunction == "SH" {
			datum := psdc.ConvertToSetCustomer(v)
			dataCustomer = append(dataCustomer, *datum)
		}
	}

	data := psdc.ConvertToSetCustomerItem(dataCustomer)

	return data
}
