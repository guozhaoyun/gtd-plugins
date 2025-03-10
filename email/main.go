package email

import (
	"github.com/gin-gonic/gin"
	"github.com/guozhaoyun/gtd-plugins/email/global"
	"github.com/guozhaoyun/gtd-plugins/email/router"
)

type emailPlugin struct {
	To       string
	From     string
	Host     string
	Secret   string
	Nickname string
	Port     int
	IsSsl    bool
}

func CreateEmailPlug(To, From, Host, Secret, Nickname string, Port int, IsSSL bool) *emailPlugin {
	global.GlobalConfig.To = To
	global.GlobalConfig.From = From
	global.GlobalConfig.Host = Host
	global.GlobalConfig.Secret = Secret
	global.GlobalConfig.Nickname = Nickname
	global.GlobalConfig.Port = Port
	global.GlobalConfig.IsSSL = IsSSL
	return &emailPlugin{}

}

func (*emailPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (*emailPlugin) RouterPath() string {
	return "email"
}
