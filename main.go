package main

import "./alfred"

//import "encoding/json"
import "fmt"
import "github.com/nlopes/slack"
import "github.com/renstrom/fuzzysearch/fuzzy"
import "os"
import "strings"

//import "net/http"
//import "strconv"

type Post struct {
	UserId int
	Id     int
	Title  string
}

func toAlfredResult(users []slack.User) string {
	result := new(alfred.Result)

	for _, user := range users {
		item := alfred.Item{
			Title:    user.Name,
			Subtitle: user.RealName,
			Arg:      user.Name,
		}

		result.Add(&item)
	}

	return result.Output()

}

func filter(users []slack.User, query string, f func(slack.User, string) bool) []slack.User {
	filtered := make([]slack.User, 0)
	for _, user := range users {
		if f(user, query) {
			filtered = append(filtered, user)
		}
	}
	return filtered
}

func main() {
	api := slack.New(os.Args[1])

	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	//users := []slack.User{
	//slack.User{
	//Name:     "b.marley",
	//RealName: "Bob Marley",
	//},
	//slack.User{
	//Name:     "j.doe",
	//RealName: "John Doe",
	//},
	//slack.User{
	//Name:     "m.saunders",
	//RealName: "Matthew Saunders",
	//},
	//}

	filterFunc := func(user slack.User, query string) bool {
		query = strings.ToLower(query)
		name := strings.ToLower(user.Name)
		realName := strings.ToLower(user.RealName)
		return fuzzy.Match(query, name) || fuzzy.Match(query, realName)
	}

	filteredUsers := filter(users, os.Args[2], filterFunc)

	fmt.Print(toAlfredResult(filteredUsers))
}
