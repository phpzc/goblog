package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, data interface{}, tplFiles ...string) {

	//加载模板
	viewDir := "resources/views/"

	for i, f := range tplFiles {
		//语法糖 将 articles.show 更正为 articles/show
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	//所有布局模板文件 slice
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	//在slice里新增我们的目标
	allFiles := append(files, tplFiles...)

	//解析模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.LogError(err)

	//渲染模板 将所有文章数据传输进去
	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}
