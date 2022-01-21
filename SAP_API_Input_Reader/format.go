package sap_api_input_reader

import (
	"sap-api-integrations-xxxxxxxx-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.XXXXXXXX
	return &requests.Header{
		XXXXXXXX:                   data.XXXXXXXX,
		ToItem: &struct {
			ToItemResults []*requests.Item `json:"results"`
		}{
			ToItemResults: []*requests.Item{
				sdc.ConvertToItem(),
			},
		},
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataHeader := sdc.XXXXXXXX
	data := sdc.XXXXXXXX.XXXXXXXXItem
	return &requests.Item{
		XXXXXXXX:                   dataHeader.XXXXXXXX,
		XXXXXXXX:                   data.XXXXXXXX,
	}
}
