package dao

import (
	"Bee-work5/model"
	"database/sql"
	"errors"
	"strconv"
)

//FindCourseById 根据课程的id查到课程的所有内容
//参数 课程的id
//返回值 该课程的所有信息 错误信息
func FindCourseById(courseId int) (model.Course,error) {
	var c model.Course
	row := DB.QueryRow("select * from course where id = ?",courseId)
	err := row.Scan(&c.Id,&c.Name,&c.Credit,&c.MaxNumber,&c.Number)
	if err==sql.ErrNoRows{
		errStr := "暂无"+strconv.Itoa(courseId)+"号课程"
		return model.Course{},errors.New(errStr)
	}else if err!=nil{
		return model.Course{},err
	}
	return c,nil
}

//UpdateWhenOperate 选课或删除课程后更新 选课人数 用户所选课程id 用户学分 用户选课数量
//参数 更新选课人数时获取tempCourse中的人数变化
//返回值 错误信息
func UpdateWhenOperate(tempC model.Course) error {
	tx,err := DB.Begin()
	if err!=nil{
		return err
	}
	sqlStr := "update course set number = ? where id = ?"   //更新选课人数
	_, err = DB.Exec(sqlStr, tempC.Number, tempC.Id)
	if err!=nil{
		_ = tx.Rollback()
		return err
	}
	sqlStr = "update user_course set class_1 = ?,class_2 = ?,class_3 = ?,class_4 = ?,class_5 = ? where account = ?"//更新用户所选的课程
	_, err = DB.Exec(sqlStr,model.CourseChosen[0],model.CourseChosen[1],model.CourseChosen[2],&model.CourseChosen[3],&model.CourseChosen[4],model.TheUser.Account)
	if err!=nil{
		_ = tx.Rollback()
		return err
	}
	sqlStr = "update user set credit = ?,sum_chosen = ? where account = ?"//更新选用户的选课总学分和选课数量
	_, err = DB.Exec(sqlStr, model.TheUser.Credit,model.TheUser.SumChosen,model.TheUser.Account)
	if err!=nil{
		_ = tx.Rollback()
		return err
	}
	_ = tx.Commit()                                             //提交事务
	return nil
}

//func AddCourse(c model.Course) error {
//	sqlStr := "insert into course (name,credit,max_number) values (?,?,?)"
//	_,err := DB.Exec(sqlStr,c.Name,c.Credit,c.MaxNumber)
//	return err
//}