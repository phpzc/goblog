package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

//Flashes flash消息数组类型 用以在会话中存储
type Flashes map[string]interface{}

//存入会话数据里的key
var flashKey = "_flashes"

func init() {
	// 在 gorilla/sessions 中存储 map 和 struct 数据需
	// 要提前注册 gob，方便后续 gob 序列化编码、解码
	gob.Register(Flashes{})
}

func Info(message string) {
	addFlash("info", message)
}

// Warning 添加 Warning 类型的消息提示
func Warning(message string) {
	addFlash("warning", message)
}

// Success 添加 Success 类型的消息提示
func Success(message string) {
	addFlash("success", message)
}

// Danger 添加 Danger 类型的消息提示
func Danger(message string) {
	addFlash("danger", message)
}

func All() Flashes {
	val := session.Get(flashKey)

	//类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}

	//读取即销毁
	session.Forget(flashKey)
	return flashMessages
}

func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}
