package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
)
func main() {
	client := &http.Client{}
	u := "hostinfo,hostid=0,hostip=127.0.0.1 cpupercent=10"
	u_byte := []byte(u)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8086/write?db=test", bytes.NewBuffer(u_byte))
	if err != nil {
		// handle error
	}

	//req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
