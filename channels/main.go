package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.amazon.com",
		"https://www.golang.org",
	}

	c := make(chan string)

	for _, link := range urls {
		go checkLink(link, c)
	}

	for _, link := range urls {
		fmt.Println(link, <-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		fmt.Println(err)
		c <- "DOWN"
		return
	}
	fmt.Println(link, "is up!")
	c <- "UP"
}
