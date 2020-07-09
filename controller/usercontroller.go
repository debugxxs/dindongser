package controller

import (
	"dingdongser/common"
	"dingdongser/models"
	"dingdongser/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service.UserService
}
//SetNewProblem 第一次登陆设置新密码和安全问题
func (uc UserController)SetNewProblem(c *gin.Context)  {
	userProblem := models.UserBindProblem{}
	if err := c.ShouldBindJSON(&userProblem);err!=nil{
		common.ErrHandler("参数有误，请检查参数",err)
		common.ResponseDataFail("参数有误，请检查参数",c)
	}
	msg, isOK :=uc.CheckUserProblem(userProblem)
	if isOK{
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}

//ModifyPass 修改密码
func (uc UserController)ModifyPass(c *gin.Context)  {
	userProblem := models.UserBindProblem{}
	if err := c.ShouldBindJSON(&userProblem);err!=nil{
		common.ErrHandler("参数有误，请检查参数",err)
		common.ResponseDataFail("参数有误，请检查参数",c)
	}
	msg, isOK :=uc.CheckUserProblemData(userProblem)
	if isOK{
		common.ResponseSuccessMsg(msg,c)
	}else {
		common.ResponseDataFail(msg,c)
	}
}
