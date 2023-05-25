package wizlib

type WikiText struct {
	Content string `json:"*"`
}

type WikiResponse struct {
	Parse struct {
		Title    string   `json:"title"`
		WikiText WikiText `json:"wikitext"`
	} `json:"parse"`
}
