package main

import "./alfred"

//import "encoding/json"
import "fmt"
import "github.com/nlopes/slack"

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

func main() {
	api := slack.New("YOUR_API_KEY")

	users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	//for _, user := range users {
	//fmt.Printf("ID: %s, Name: %s, RealName: %s\n", user.ID, user.Name, user.RealName)
	//}

	//r, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	//if err != nil {
	//panic(err.Error())
	//}

	//defer r.Body.Close()

	fmt.Print(toAlfredResult(users))
}
