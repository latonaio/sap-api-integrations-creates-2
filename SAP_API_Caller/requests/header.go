package requests

type Header struct {
	XXXXXXXX             string        `json:"XXXXXXXX"`
	XXXXXXXX            *string        `json:"XXXXXXXX"`
	ToItem              *ToItem        `json:"to_Item"`
}

type ToItemResults struct {
	XXXXXXXX             string  `json:"XXXXXXXX"`
	XXXXXXXX            *string  `json:"XXXXXXXX"`
}

type ToItem struct {
	ToItemResults []ToItemResults `json:"results"`
}
