package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "seems to be down")
		c <- link
		return
	}

	fmt.Println(link, "looks fine")
	c <- link
}

func main() {
	links := []string{
		"http://facebook.com",
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	for l := range c {

		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
		//checkLink(l, c)
	}
}
