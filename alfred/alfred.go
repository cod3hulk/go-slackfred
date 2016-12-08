package alfred

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

type Result struct {
	Items []Item `json:"items"`
}

func (r *Result) Add(item *Item) *Result {
	r.Items = append(r.Items, *item)

	return r
}

func (r *Result) Output() string {
	output, err := json.Marshal(*r)

	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", output)
}
