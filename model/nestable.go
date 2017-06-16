package model

import (
	"encoding/json"
	"fmt"
)

type Handle struct {
	Id string `json:"id"`
}

type Children struct {
	Id 	  string 	`json:"id"`
	Children []Handle 	`json:"children"`
}

type Nestable struct {
	Name 	string `json:"name"`
	Group []Children `json:"group"`
}

func Test() *Nestable {
	body := `
	{
		"name":"lxp",
		"group":[
			{
				"id":"sn",
				"children":[{"id":"3"},{"id":"3"},{"id":"3"},{"id":"3"},{"id":"3"}]
			},{
				"id":"jdbc",
				"children":[{"id":"3"},{"id":"3"},{"id":"3"}]
			}
			]
	}`
	//body := `{"name":"lxp","group":[{"id":1},{"id":2}]}`
	var r Nestable
	err := json.Unmarshal([]byte(body),&r)
	if err != nil {
		fmt.Printf("err was %v",err)
	}
	//fmt.Println(r.Group[0].Handle.Id)
	//fmt.Println(r.Group[0].Id)
	return &r
}

func TestStruct() *Nestable {
	var JsonResult Nestable
	JsonResult.Name = "Sn"

	JsonChildren := []Children{}
	JsonChildren = append(JsonChildren,Children{
		Id:"主机信息",
		Children:[]Handle{
			Handle{Id:"hostname"},
			Handle{Id:"idc"},
		},
	})

	JsonResult.Group = JsonChildren
	b,err := json.Marshal(&JsonResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("1111111111111",string(b))
	return &JsonResult
}