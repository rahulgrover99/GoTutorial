package main

import (
	"fmt"
	"net/http"
	"time"
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

	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down.")
		fmt.Println(err)
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
