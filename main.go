package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"crypto/tls"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 2 {
		fmt.Println("Too few arguments, need url and number repeats")
		return
	}
	url := argsWithoutProg[0]
	count, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		fmt.Println(err)
	}

	testUrl(url, count)

}

func testUrl(url string, count int) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}
	for i := 0; i < count; i++ {
		client := &http.Client{Transport: tr}
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(body))
		resp.Body.Close()
	}
}
