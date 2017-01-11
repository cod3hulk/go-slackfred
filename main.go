package main

import "./alfred"

import "fmt"
import "github.com/nlopes/slack"
import "github.com/renstrom/fuzzysearch/fuzzy"
import "os"
import "strings"

func usersToAlfredItems(users []slack.User) []alfred.Item {
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

func groupsToAlfredItems(groups []slack.Group) []alfred.Item {
	items := make([]alfred.Item, 0)

	for _, group := range groups {
		if !strings.HasPrefix(group.Name, "mpdm") {
			item := alfred.Item{
				Title:    group.Name,
				Subtitle: "Group",
				Arg:      group.Name,
			}
			items = append(items, item)
		}
	}

	return items
}

func channelsToAlfredItems(channels []slack.Channel) []alfred.Item {
	items := make([]alfred.Item, 0)

	for _, channel := range channels {
		item := alfred.Item{
			Title:    channel.Name,
			Subtitle: "Channel",
			Arg:      channel.Name,
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

	channels, err := api.GetChannels(true)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	groups, err := api.GetGroups(true)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	result := new(alfred.Result).
		AddAll(usersToAlfredItems(users)).
		AddAll(channelsToAlfredItems(channels)).
		AddAll(groupsToAlfredItems(groups)).
		Filter(os.Args[2], func(item alfred.Item, query string) bool {
			query = strings.ToLower(query)
			arg := strings.ToLower(item.Arg)
			subtitle := strings.ToLower(item.Subtitle)
			return fuzzy.Match(query, arg) || fuzzy.Match(query, subtitle)
		}).Output()

	fmt.Print(result)
}
