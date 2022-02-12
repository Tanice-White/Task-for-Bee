package api

import (
	"Bee-work5/model"
	"Bee-work5/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func TOShowCourse(c *gin.Context)  {
	c.JSON(200,gin.H{
		"课程总数":model.TheUser.SumChosen,
		"已选学分":model.TheUser.Credit,
	})
	for k,v := range model.CourseName {
		c.JSON(200,gin.H{
			"course id":k+1,
			"course name--credit":v,
		})
	}
}

func TOShowCourseChosen(c *gin.Context)  {
	c.JSON(200,gin.H{
		"课程总数":model.TheUser.SumChosen,
		"已选学分":model.TheUser.Credit,
	})
	for k,v := range model.CourseName {
		for _,n := range model.CourseChosen{
			if k+1==n{
				c.JSON(200,gin.H{
					"course id":k+1,
					"course name--credit":v,
				})
			}
		}
	}
}

func ToChooseCourse(c *gin.Context) {
	if model.TheUser.SumChosen == model.MaxSumChosen {      //选修课程总数最多5门
		c.JSON(200, gin.H{"msg": "选修课最多选择5门"})
		return
	}
	if model.TheUser.Credit == model.MaxCredit {            //最多选修10分
		c.JSON(200, gin.H{"msg": "选修学分最多10分"})
		return
	}
	id, err := strconv.Atoi(c.Query("id"))             //读取选项
	if err != nil {
		c.JSON(200, gin.H{
			"code":501,
			"err": err.Error()})
		return
	}
	err = service.ChooseCourse(id)                          //记录选择课程
	if err != nil {
		c.JSON(200, gin.H{
			"code":501,
			"err": err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{"msg":"选课成功"})
}

func ToDeleteCourse(c *gin.Context){
	if model.TheUser.SumChosen == 0 {
		c.JSON(200, gin.H{"msg": "您还未选择课程"})
		return
	}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"code":501,
			"err": err.Error()})
		return
	}
	err = service.DeleteCourse(id)      //删除已选的课程
	if err != nil {
		c.JSON(200, gin.H{
			"code":501,
			"err": err.Error()})
		return
	}
	c.JSON(200,gin.H{"msg":"删除成功"})
}