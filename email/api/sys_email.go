package api

import (
	"github.com/gin-gonic/gin"
	email_response "github.com/guozhaoyun/gtd-plugins/email/model/response"
	"github.com/guozhaoyun/gtd-plugins/email/service"
	"github.com/guozhaoyun/gtd2110/server/global"
	"github.com/guozhaoyun/gtd2110/server/model/common/response"
	"go.uber.org/zap"
)

type EmailApi struct {
}

// @Tags System
// @Summary 发送测试邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/emailTest [post]
func (s *EmailApi) EmailTest(c *gin.Context) {
	if err := service.ServiceGroupApp.EmailTest(); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}

// @Tags System
// @Summary 发送邮件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body email_response.Email true "发送邮件必须的参数"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发送成功"}"
// @Router /email/sendEmail [post]
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email email_response.Email
	_ = c.ShouldBindJSON(&email)
	if err := service.ServiceGroupApp.SendEmail(email.To, email.Subject, email.Body); err != nil {
		global.GVA_LOG.Error("发送失败!", zap.Any("err", err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
