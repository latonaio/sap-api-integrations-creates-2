package requests

type Header struct {
	XXXXXXXX             string        `json:"XXXXXXXX"`
	XXXXXXXX            *string        `json:"XXXXXXXX"`
	ToItem              *struct {
		ToItemResults []*Item `json:"results"`
	} `json:"to_Item"`
}
