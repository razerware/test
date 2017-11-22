package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
)

type user struct {
	Id    int         `json:"-"`
	Name  interface{} `json:"username"`
	Age   int         `json:"age"`
	Email string      `json:"email"`
}

func main() {
	client := &http.Client{}
	u := user{Id: 200, Name: "lzy", Age: 10, Email: "lizhengyin1115@163.com"}
	u_json, _ := json.Marshal(u)
	u_byte := []byte(u_json)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/lzy", bytes.NewBuffer(u_byte))
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
