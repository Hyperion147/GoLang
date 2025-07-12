package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const webUrl = "https://todo.suryansu.pro/auth?hello=hyper"

func main() {
	topic := "This program is to learn web requests."
	fmt.Println(topic)

	res, err := http.Get(webUrl)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	result, _ := url.Parse(webUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.RawQuery)

	fmt.Printf("Response: %T\n", res)
	fmt.Println("Response: ")

	data, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	constructUrl := &url.URL{
		Scheme: "https",
		Host:   "suryansu.pro",
	}

	constructedUrl := constructUrl.String()

	fmt.Println(constructUrl)
	fmt.Println(constructedUrl)

}
