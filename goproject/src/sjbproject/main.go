package main

import (
	"github.com/astaxie/beego"
	_ "sjbproject/models"
	_ "sjbproject/routers"
)

func main() {
	beego.AddFuncMap("showPrePage", HandlePrePage)
	beego.AddFuncMap("showNextPage", HandleNextPage)
	beego.AddFuncMap("isSelect", HandleIsSelect)
	beego.Run()
}

//计算上一页
func HandlePrePage(pageIndex int) (result int) {
	result = pageIndex - 1
	if result <= 0 {
		result = 1
	}
	return
}

//计算下一页
func HandleNextPage(pageIndex int) (result int) {
	result = pageIndex + 1
	return
}

//判断显示默认的文章类型
func HandleIsSelect(typeName string, currentTypeName string) (isSelect bool) {

	if typeName == currentTypeName {
		return true
	} else {
		return false
	}
}
