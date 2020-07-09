package models

type Userproblem struct {
	ProblemId	int64	`json:"problemId" xorm:"pk autoincr problemId"`
	Problem 	string	`json:"problem" xorm:"varchar(64) problem"`
	Answer	string	`json:"answer" xorm:"varchar(64) answer"`
}

type ProblemUserModel struct {
	User	`xorm:"extends"`
	Userproblem `xorm:"extends"`
}
//UserBindProblem 第一次登陆空密码需要用到的空实体
type UserBindProblem struct {
	UserName string	`json:"userName"`
	Problem string	`json:"problem"`
	Password string	`json:"password"`
	Answer string	`json:"answer"`
}
