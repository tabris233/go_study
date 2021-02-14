package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	rsp  = make(chan string, 1)
	done = make(chan struct{})
)

func fetch(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		// os.Exit(1)
		return "+++"
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		return "---"
	}
	// fmt.Printf("%s", b)
	return string(b)
}

func getChannel(url string) {
	for {
		select {
		case rsp <- fetch(url):
			return
		case <-done:
			fmt.Println(url, "end")
			for range rsp {
			}
			return
		}
	}
}

func search() string {
	go getChannel("http://bing.com")
	go getChannel("http://google.com")
	go getChannel("http://baidu.com")
	defer close(done)
	return <-rsp
}

func main() {
	result := search()

	fmt.Println(result)

	time.Sleep(time.Second * 10)
}
