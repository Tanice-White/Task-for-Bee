package service

import (
	"Bee-work5/dao"
	"Bee-work5/model"
	"errors"
)

//ChooseCourse 在本地储存中增加所选的课程，暂不写入数据库，且不允许重复
//参数 所选课程的id
//返回值 错误信息
func ChooseCourse(courseId int)error{
	for _,v:= range model.CourseChosen{
		if v==courseId{
			return errors.New("您已选此课程")
		}
	}
	tempC,err := dao.FindCourseById(courseId)
	if err!=nil{
		return err
	}
	if tempC.Number==tempC.MaxNumber{            //课程人数已满的情况
		return errors.New("课程人数已满")
	}
	if (model.TheUser.Credit+tempC.Credit)>model.MaxCredit{//学分超出上限
		return errors.New("选课学分之和不能超过10分")
	}
	flag := false                               //表示是否找到了该课程
	for k,v := range model.CourseChosen{
		if v==-1{
			model.CourseChosen[k]=courseId       //选课成功
			model.TheUser.SumChosen += 1         //所选课程总数+1
			model.TheUser.Credit += tempC.Credit //加上学分
			tempC.Number++                       //选课的学生增加
			flag = true
			break                                //找到一个空缺即可
		}
	}
	if !flag {                                   //没有找到此课程
		return errors.New("您还未选择此课程")
	}
	err = dao.UpdateWhenOperate(tempC)
	return err
}

//DeleteCourse 在本地储存中删除所选的课程，暂不写入数据库
//参数 所选课程的id
//返回值 错误信息
func DeleteCourse(courseId int) error {
	tempC,err := dao.FindCourseById(courseId)
	if err!=nil{
		return err
	}
	flag := false                               //表示是否找到了该课程
	for k,v := range model.CourseChosen{
		if v==courseId {
			model.CourseChosen[k] = -1          //删除课程
			model.TheUser.SumChosen -= 1        //已选课程数量-1
			model.TheUser.Credit -= tempC.Credit//减去删除课程的学分
			tempC.Number--                      //选课学生的数量-1
			flag = true
			break
		}
	}
	if !flag {                                //没有找到此课程
		return errors.New("您还未选择此课程")
	}
	err = dao.UpdateWhenOperate(tempC)
	return err
}