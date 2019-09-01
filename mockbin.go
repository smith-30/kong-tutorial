package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {

	url := "http://mockbin.com/request?foo=bar&foo=baz"

	payload := strings.NewReader("{\"foo\": \"bar\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cookie", "foo=bar; bar=baz")
	req.Header.Add("x-pretty-print", "2")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	pp.Print(res)
	fmt.Println(string(body))

}
