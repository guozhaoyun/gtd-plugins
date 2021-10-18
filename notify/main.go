package notify

import (
	"github.com/guozhaoyun/gtd-plugins/notify/global"
)

type ddPlugin struct {
	Secret string
	Token  string
	Url    string
}

func CreateDDPlug(url, Token, Secret string) *ddPlugin {
	global.GlobalConfig_.Url = url
	global.GlobalConfig_.Token = Token
	global.GlobalConfig_.Secret = Secret
	return &ddPlugin{}
}

func (*ddPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitRouter(group)
}

func (*ddPlugin) RouterPath() string {
	return "notify"
}
