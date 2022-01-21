package responses

type Header struct {
	D struct {
		XXXXXXXX            string `json:"XXXXXXXX"`
		ToItem              struct {
			ToItemResults []Item `json:"results"`
		} `json:"to_Item"`
	} `json:"d"`
}
