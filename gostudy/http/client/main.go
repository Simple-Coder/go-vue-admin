package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := new(http.Client)
	request, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	res, _ := client.Do(request)
	all, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(all))
}
