package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guozhaoyun/gtd-plugins/notify/api"
	"github.com/guozhaoyun/gtd2110/server/middleware"
)

type NotifyRouter struct {
}

func (s *NotifyRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Use(middleware.OperationRecord())
	var SendTextMessage = api.ApiGroupApp.Api.SendTextMessage
	{
		router.POST("sendTextMessage", SendTextMessage)
	}
}
