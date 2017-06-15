# 需求

* golang将数据转换成json格式

'''
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
'''

* 将json转换成nestable插件需要的html代码

* 通过net/http 生成