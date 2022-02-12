package model

type UserCourse struct {
	Id int64        `json:"id"`
	Name string     `json:"name"`
	Account string  `json:"account"`
	Class1 int32    `json:"class_1"`
	Class2 int32    `json:"class_2"`
	Class3 int32    `json:"class_3"`
	Class4 int32    `json:"class_4"`
	Class5 int32    `json:"class_5"`
}

var CourseChosen = [5]int{-1,-1,-1,-1,-1}    //记录已选课程