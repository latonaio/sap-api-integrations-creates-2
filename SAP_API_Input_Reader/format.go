package sap_api_input_reader

import (
	"sap-api-integrations-xxxxxxxx-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	return &requests.Header{
		XXXXXXXX:                   sdc.XXXXXXXX.XXXXXXXX,
		ToItem:                     sdc.ConvertToToItem(),
	}
}

func (sdc *SDC) ConvertToToItem() *requests.ToItem {
	return &requests.ToItem{
		ToItemResults: []requests.ToItemResults{
			{
				XXXXXXXX:           sdc.XXXXXXXX.XXXXXXXX,
			},
		},
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	return &requests.Item{
		XXXXXXXX:                   sdc.XXXXXXXX.XXXXXXXX,
	}
}
