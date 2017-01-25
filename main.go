package main

import "github.com/cod3hulk/go-slackfred/alfred"

import "fmt"
import "github.com/nlopes/slack"
import "github.com/renstrom/fuzzysearch/fuzzy"
import "os"
import "strings"
import "sync"

func users(apiToken string, items chan alfred.Item, wg *sync.WaitGroup) {
	defer wg.Done()
	api := slack.New(apiToken)

	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, user := range users {
		items <- alfred.Item{
			Title:    user.Name,
			Subtitle: user.RealName,
			Arg:      user.Name,
		}
	}

}

func channels(apiToken string, items chan alfred.Item, wg *sync.WaitGroup) {
	defer wg.Done()
	api := slack.New(apiToken)

	channels, err := api.GetChannels(true)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, channel := range channels {
		items <- alfred.Item{
			Title:    channel.Name,
			Subtitle: "Channel",
			Arg:      channel.Name,
		}
	}

}

func groups(apiToken string, items chan alfred.Item, wg *sync.WaitGroup) {
	defer wg.Done()
	api := slack.New(apiToken)

	groups, err := api.GetGroups(true)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	for _, group := range groups {
		if !strings.HasPrefix(group.Name, "mpdm") {
			items <- alfred.Item{
				Title:    group.Name,
				Subtitle: "Group",
				Arg:      group.Name,
			}
		}
	}

}

func filterItem(item alfred.Item, query string) bool {
	query = strings.ToLower(query)
	arg := strings.ToLower(item.Arg)
	subtitle := strings.ToLower(item.Subtitle)
	return fuzzy.Match(query, arg) || fuzzy.Match(query, subtitle)
}

func main() {
	items := make(chan alfred.Item)
	var wg sync.WaitGroup

	result := new(alfred.Result)

	go func(items chan alfred.Item, result *alfred.Result) {
		for item := range items {
			result.Add(&item)
		}
	}(items, result)

	wg.Add(3)
	go users(os.Args[1], items, &wg)
	go groups(os.Args[1], items, &wg)
	go channels(os.Args[1], items, &wg)

	wg.Wait()
	close(items)

	fmt.Print(result.Filter(os.Args[2], filterItem).Output())
}
