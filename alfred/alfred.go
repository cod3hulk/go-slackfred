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

func (r *Result) AddAll(items []Item) *Result {
	r.Items = append(r.Items, items...)

	return r
}

func (r *Result) AddChannel(items chan Item, query string, f func(Item, string) bool) *Result {
	for item := range items {
		if f(item, query) {
			r.Items = append(r.Items, item)
		}
	}

	return r
}

func (r *Result) Filter(query string, f func(Item, string) bool) *Result {
	items := make([]Item, 0)
	for _, item := range r.Items {
		if f(item, query) {
			items = append(items, item)
		}
	}

	r.Items = items

	return r
}

func (r *Result) Output() string {
	output, err := json.Marshal(*r)

	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", output)
}
