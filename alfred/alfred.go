package alfred

type Item struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

type Result struct {
	Items []Item `json:"items"`
}
