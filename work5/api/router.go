package api

import (
	"github.com/gin-gonic/gin"
)

//StartEngine 启动相关引擎并且挂载所有的入口
func StartEngine() {
	engine := gin.Default()

	engine.POST("/register",ToRegister)
	engine.GET("/login",CheckAuthorizationBeforeLogIn,ToLogIn,GenerateToken)
	engine.GET("/help",GetHelp)

	engineGroup := engine.Group("/user",CheckAuthorization)
	{
		engineGroup.GET("/choose",TOShowCourse,ToChooseCourse,TOShowCourse)
		engineGroup.GET("/delete",TOShowCourseChosen,ToDeleteCourse,TOShowCourseChosen)
		engineGroup.GET("/change",ToChangePassword)
	}
	_ = engine.Run(":9090")
}