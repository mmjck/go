package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func mainn() {
	resp, err := http.Get("http://localhost:8080/")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Not success", resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)

	var response []User
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println("Erro ao recuperar dados", err.Error())
		return
	}

	fmt.Println(response)
}
