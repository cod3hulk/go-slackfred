package main

import "./alfred"

import "fmt"
import "github.com/nlopes/slack"
import "github.com/renstrom/fuzzysearch/fuzzy"
import "os"
import "strings"

func toAlfredItems(users []slack.User) []alfred.Item {
	items := make([]alfred.Item, 0)

	for _, user := range users {
		item := alfred.Item{
			Title:    user.Name,
			Subtitle: user.RealName,
			Arg:      user.Name,
		}
		items = append(items, item)
	}

	return items
}

func main() {
	api := slack.New(os.Args[1])

	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	result := new(alfred.Result).AddAll(toAlfredItems(users)).Filter(os.Args[2], func(item alfred.Item, query string) bool {
		query = strings.ToLower(query)
		arg := strings.ToLower(item.Arg)
		subtitle := strings.ToLower(item.Subtitle)
		return fuzzy.Match(query, arg) || fuzzy.Match(query, subtitle)
	}).Output()

	fmt.Print(result)
}
