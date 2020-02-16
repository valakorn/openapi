package main

import (
	"encoding/json"
	"fmt"
)

var data = `[{
	"userId": 1,
	"id": 1,
	"title": "delectus aut autem",
	"completed": false
  }]`

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	dataStruct := []Todo{}
	v := &dataStruct
	fmt.Println(dataStruct)
	json.Unmarshal([]byte(data), v)
	fmt.Println(dataStruct)
	//===========================================
	fmt.Println(len(dataStruct)) //Count Data Json

	dataStruct[0].Completed = true
	dataStruct[1].Completed = true

}
