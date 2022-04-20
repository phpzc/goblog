package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, name string, data interface{}) {

	//加载模板
	viewDir := "resources/views/"

	//语法糖 将 articles.show 更正为 articles/show
	name = strings.Replace(name, ".", "/", -1)

	//所有布局模板文件 slice
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	//在slice里新增我们的目标
	newFiles := append(files, viewDir+name+".gohtml")

	//解析模板文件
	tmpl, err := template.New(name + ".gohtml").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(newFiles...)
	logger.LogError(err)

	//渲染模板 将所有文章数据传输进去
	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
