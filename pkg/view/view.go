package view

import (
	"embed"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"io/fs"
	"strings"
)

//定义通用 传参给视图的数据
type D map[string]interface{}

var TplFS embed.FS

func Render(w io.Writer, data D, tplFiles ...string) {

	RenderTemplate(w, "app", data, tplFiles...)
}

func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string) {

	//通用模板数据
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User()
	data["flash"] = flash.All()
	data["Users"], _ = user.All()
	data["Categories"], _ = category.All()

	//生成模板文件
	allFiles := getTemplateFiles(tplFiles...)

	//解析模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).ParseFS(TplFS, allFiles...)
	logger.LogError(err)

	//渲染模板 将所有文章数据传输进去
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}

func getTemplateFiles(tplFiles ...string) []string {
	//加载模板
	viewDir := "resources/views/"

	//遍历传参文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tplFiles {

		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	//所有布局模板文件 slice
	layoutFiles, err := fs.Glob(TplFS, viewDir+"layouts/*.gohtml")
	logger.LogError(err)

	//合并所有文件
	return append(layoutFiles, tplFiles...)
}
