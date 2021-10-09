package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const myIpServerUrl = "http://ip.modsaid.com"

func main(){
	resp, err := http.Get(myIpServerUrl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ip := string(bytes)
	fmt.Println(ip)
}
