package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	ioutil.WriteFile("page.html", bs, 0666)

}
