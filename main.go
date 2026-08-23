package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://crt.name/v1/search?apex=x.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var urls []string
	for i := 1; scanner.Scan(); i++ {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	check_urls(urls[:10])
	// time.Sleep(time.Second * 60)

}

func check_urls(urls []string) {
	result := make(chan string, len(urls))
	for _, u := range urls {
		go ping(u, result)
	}
	for msg := range result {
		fmt.Println(msg)
	}
	// fatal error: all goroutines are asleep - deadlock!
}

func ping(s string, done chan string) {
	u := fmt.Sprintf("https://%s", s)
	resp, err := http.Get(u)
	if err != nil {
		u = fmt.Sprintf("http://%s", s)
		resp, err = http.Get(u)
		if err != nil {
			done <- fmt.Sprintf(s, "500")
			return
		}
	}
	defer resp.Body.Close()

	done <- fmt.Sprintf(s, resp.Status)
}
