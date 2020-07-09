package service

import (
	"dingdongser/dao"
	"dingdongser/models"
)

type UserService struct {
	dao.UserDao
}

//CheckUserPass 用户登陆相关操作
func (us UserService)CheckUserPass(userName,userPass string)bool  {
	//将密码解码
	user := us.GetUserPass(userName)
	if user.UserName == userName &&user.PassWord == userPass{
		return true
	}
	return false
}
//CheckUserProblem 设置安全问题和新密码
func (us UserService)CheckUserProblem(problem models.UserBindProblem)(string,bool)  {
	userPass := problem.Password
	userName := problem.UserName
	answer :=problem.Answer
	problemStr := problem.Problem
	//设置安全问题
	msg,proRes:=us.SaveProblem(userName,problemStr,answer)
	if proRes >0{
		passMsg,passRes:=us.SaveUserPass(userName,userPass)
		if passRes >0{
			return passMsg,true
		}
	}
	return msg,false
}
//CheckUserProblemData 根据安全问题修改密码
func (us UserService)CheckUserProblemData(problem models.UserBindProblem)(string,bool){
	userPass := problem.Password
	userName := problem.UserName
	answer :=problem.Answer
	problemStr := problem.Problem

	//查询安全问题
	ProMsg,isExit:=us.QueryProblem(problemStr,answer)
	if isExit{
		//可以修改密码
		passMsg,res:=us.SaveUserPass(userName,userPass)
		if res >0{
			return passMsg,true
		}
		return passMsg,false
	}
	return ProMsg,false
}
