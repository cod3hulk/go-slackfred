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

func encode(posts []Post) alfred.Result {
	items := []alfred.Item{}

	for _, v := range posts {
		items = append(items, alfred.Item{strconv.Itoa(v.Id), v.Title, strconv.Itoa(v.UserId)})
	}

	return alfred.Result{Items: items}

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

	test := encode(res)

	b, err := json.Marshal(test)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)
}
