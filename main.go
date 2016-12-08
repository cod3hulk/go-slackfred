// Package main provides ...
package main

import "./alfred"
import "encoding/json"
import "fmt"
import "net/http"
import "strconv"

type Post struct {
	UserId int
	Id     int
	Title  string
}

func toAlfredResult(posts []Post) string {
	result := new(alfred.Result)

	for _, v := range posts {
		item := alfred.Item{
			Title:    strconv.Itoa(v.Id),
			Subtitle: v.Title,
			Arg:      strconv.Itoa(v.UserId),
		}

		result.Add(&item)
	}

	return result.Output()

}

func main() {
	r, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err.Error())
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	res := make([]Post, 0)

	err = decoder.Decode(&res)
	if err != nil {
		panic(err.Error())
	}

	fmt.Print(toAlfredResult(res))
}
