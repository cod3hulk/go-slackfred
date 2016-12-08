// Package main provides ...
package main

import "./alfred"
import "encoding/json"
import "fmt"
import "net/http"
import "os"
import "strconv"

type Post struct {
	UserId int
	Id     int
	Title  string
}

func encode(posts []Post) *alfred.Result {
	result := new(alfred.Result)

	for _, v := range posts {
		result.Add(alfred.Item{
			Id:     strconv.Itoa(v.Id),
			Title:  v.Title,
			UserId: strconv.Itoa(v.UserId),
		})
	}

	return result

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

	result := encode(res)

	fmt.Print(result)
}
