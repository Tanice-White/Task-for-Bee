package model

//User 储存学生信息
//包括 姓名 账号 密码 已选课程的总学分 所选课程的id
type User struct {
	Id int64          `json:"id"`
	Name string       `json:"name"`
	Account string    `json:"account"`
	Password string   `json:"password"`
	Credit int32      `json:"credit"`
	SumChosen int32   `json:"sum_chosen"`
}

//TempUser 登录和注册时获取账号和密码的结构
type TempUser struct {
	Name string       `json:"name"`
	Account string    `json:"account"`
	Password string   `json:"password"`
}
var TheUser User              //储存当前登录者的信息
var Sum int64                 //记录总人数
var MaxSumChosen int32 = 5    //选课最多5门
var MaxCredit int32 = 10      //最多修10分

