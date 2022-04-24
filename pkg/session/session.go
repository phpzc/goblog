package session

import (
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/sessions"
)

//Store gorilla sessions 的存储库
var Store = sessions.NewCookieStore([]byte("33446a9dcf9ea060a0a6532b166da32f304af0de"))

//Session 当前会话
var Session *sessions.Session

//Request 用以获取会话
var Request *http.Request

//Response 用以写入会话
var Response http.ResponseWriter

//StartSession 初始化会话 在中间件中调用
func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error

	Session, err = Store.Get(r, "goblog-session")
	logger.LogError(err)

	Request = r
	Response = w

}

func Put(key string, value interface{}) {

	Session.Values[key] = value
	Save()
}

func Get(key string) interface{} {
	return Session.Values[key]
}

func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

//删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

func Save() {
	err := Session.Save(Request, Response)
	logger.LogError(err)
}