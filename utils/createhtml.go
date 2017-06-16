package utils

import (
	"fmt"
	"strings"
	"github.com/lflxp/json/model"
)

type CreateHtml interface {
	SetHandle(data string) 		string
	SetItem(id int,handle string)	string
	SetItemStart(id int,handle string) string
	SetItemEnd() 			string
	SetList(items []string) 	string
	Parse(data *model.Nestable) 	string
}

func SetHandle(data string) string {
	return fmt.Sprintf(`<div class="dd-handle">%s</div>`,data)
}

func SetItem(id int,handle string) string {
	return fmt.Sprintf(`<li class="dd-item" data-id="%d">%s</li>`,id,SetHandle(handle))
}

func SetItemStart(id int,handle string) string {
	return fmt.Sprintf(`<li class="dd-item" data-id="%d">%s`,id,SetHandle(handle))
}

func SetItemEnd() string {
	return "</li>"
}

func SetList(items []string) string {
	return fmt.Sprintf(`<ol class="dd-list">%s</ol>`,strings.Join(items,""))
}

/**
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
	}
 */
func Parse(data *model.Nestable) string {
	//生成name的菜单
	menu := SetItem(0,data.Name)
	//整体数据
	tmp := []string{}
	for x,y := range data.Group {
		//第一层 数据集
		list_one := []string{}
		//children层 数据集 涉及到OL的开(<ol>)和闭（</ol>）
		list_two := []string{}
		//生成OL开 无/ol闭
		list_one = append(list_one,SetItemStart(x,y.Id))
		for n,m := range y.Children {
			//生成Children ol items数据
			list_two = append(list_two,SetItem(n,m.Id))
		}
		//生成children ol数据字符集
		list_one = append(list_one,SetList(list_two))
		//ol闭 </ol>
		list_one = append(list_one,SetItemEnd())
		//将一个group的数据加入整体数据
		tmp = append(tmp,list_one...)
	}
	fmt.Println(menu+strings.Join(tmp,"\n"))
	//第一层加OL 闭合
	return "<ol class=\"dd-list\">"+menu+strings.Join(tmp,"")+"</ol>"
}

func TestHtml() string {
	data := model.Test()
	return Parse(data)
}