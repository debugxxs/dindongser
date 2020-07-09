package dao

import (
	"dingdongser/common"
	models "dingdongser/models"
	"dingdongser/tools"
	"fmt"
)

type UserDao struct {

}


//GetUserId 获取用户id
func (ud UserDao)GetUser(userName string)models.User {
	user := models.User{}
	_, _ = tools.DbEngine.Table("user").Where("userName=?", userName).Get(&user)
	return user
}


//GetProblemId 获取最新修改安全问题id
func (ud UserDao)GetProblemId()int64{
	var problemList []models.Userproblem
	_ = tools.DbEngine.Table("userproblem").Desc("problem").Find(&problemList)
	return problemList[0].ProblemId
}

//获取最新插入的密码id
func (ud UserDao)GetPassId()int64 {
	var passwordList []models.Password
	_ = tools.DbEngine.Table("password").Desc("passWord").Find(&passwordList)
	return passwordList[0].PassId
}

//GetUserPass 获取用户登陆密码
func (ud UserDao)GetUserPass(userName string)models.UserPass  {
	userpass := models.UserPass{}
	_, _ = tools.DbEngine.Table("user").Join("INNER", "password", "user.passId = password.passId and user.userName = ?", userName).Get(&userpass)
	return userpass
}


//saveProblem 保存问题到数据库
func (ud UserDao)SaveProblem(userName,problemStr,answer string) (string,int64) {
	//把用户问题保存到数据库
	//1先根据用户表中的id来查询用户原来的problemId
	//2再根据problemId的值来修改，如果值为0就新增数据并更新user表
	//3.如果值为非0就删除原来表中的对于id数据，并插入新数据，并重新更新user表
	problem := models.Userproblem{Problem: problemStr,Answer: answer}
	user := ud.GetUser(userName)
	if user.ProblemId ==0{
		res, _ := tools.DbEngine.Table("userproblem").InsertOne(&problem)
		if res >0{
			//插入成功更新到user表
			problemId := ud.GetProblemId()
			newUser := models.User{ProblemId: problemId}
			userRes, _ := tools.DbEngine.Table(user).ID(user.UserId).Update(&newUser)
			if userRes >0{
				return common.UpDataSuccess,userRes
			}
			return common.UpDataFail,res
		}else {
			return common.InsertDataFail,res
		}
	}else {
		userPro := new(models.Userproblem)
		_, _ = tools.DbEngine.Table("userproblem").ID(user.ProblemId).Delete(userPro)
		res, _ := tools.DbEngine.Table("userproblem").InsertOne(&problem)
		if res>0{
			problemId := ud.GetProblemId()
			fmt.Println(problemId)
			newUser := models.User{ProblemId: problemId}
			userRes, _ := tools.DbEngine.Table(user).ID(user.UserId).Update(&newUser)
			if userRes >0{
				return common.UpDataSuccess,userRes
			}
			return common.UpDataFail,userRes
		}
		return common.InsertDataFail,res
	}
}


//SaveUserPass 保存密码到数据库
func (ud UserDao)SaveUserPass(userName ,userPass string)(string,int64){
	//把用户问题保存到数据库
	//1先根据用户表中的id来查询用户原来的passId
	//2再根据passId的值来修改，如果值为0就新增数据并更新user表
	//3.如果值为非0就删除原来表中的对于id数据，并插入新数据，并重新更新user表
	password := models.Password{PassWord: userPass}
	user := ud.GetUser(userName)
	if user.PassId !=1{
		//密码不属于默认密码，需要删除密码表中的旧密码再新插入数据
		pass :=new(models.Password)
		res, _ := tools.DbEngine.Table("password").ID(user.PassId).Delete(pass)
		res, _ = tools.DbEngine.Table("password").InsertOne(&password)
		if res>0{
			//把最新的数据更新到用户表中
			passId := ud.GetPassId()
			newUser := models.User{PassId: passId}
			userRes, _ := tools.DbEngine.Table("user").ID(user.UserId).Update(&newUser)
			if userRes>0{
				return common.UpDataSuccess,userRes
			}
			return common.UpDataFail,res
		}
		return common.InsertDataFail,res
	}else {
		res, _ := tools.DbEngine.Table("password").InsertOne(&password)
		if res>0{
			passId := ud.GetPassId()
			newUser := models.User{PassId: passId}
			userRes, _ := tools.DbEngine.Table("user").ID(user.UserId).Update(&newUser)
			if userRes>0{
				return common.UpDataSuccess,userRes
			}
			return common.UpDataFail,res
		}
		return common.InsertDataFail,res
	}
}

//QueryProblem 查询安全问题和答案
func (ud UserDao)QueryProblem(problemStr,answer string)(string,bool)  {
	problem:=models.Userproblem{}
	res,_:=tools.DbEngine.Table("problem").Where("problem = ? and answer = ?",problemStr,answer).Get(&problem)
	if res{
		return common.QueryDataSuccess,true
	}
	return common.UpDataFail,false
}
