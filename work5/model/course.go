package model

type Course struct {
	Id int                `json:"id"`
	Name string           `json:"name"`
	Credit int32          `json:"credit"`
	MaxNumber int32       `json:"max_number"`
	Number int32          `json:"number"`
}

//CourseName 作用于测试
var CourseName = map[int]string{
	0:"高等数学--4",
	1:"体育--2",
	2:"文言文鉴赏--2",
	3:"C/C++编程--4",
	4:"英语--4",
	5:"论文写作指导--1",
}
//var C1 = Course{
//	Name: "高等数学",
//	Credit: 4,
//	MaxNumber: 100,
//}
//var C2 = Course{
//	Name: "体育",
//	Credit: 2,
//	MaxNumber: 35,
//}
//var C3 = Course{
//	Name: "文言文鉴赏",
//	Credit: 2,
//	MaxNumber: 120,
//}
//var C4 = Course{
//	Name: "C/C++编程",
//	Credit: 4,
//	MaxNumber: 30,
//}
//var C5 = Course{
//	Name: "英语",
//	Credit: 4,
//	MaxNumber: 60,
//}
//var C6 = Course{
//	Name: "论文写作指导",
//	Credit: 1,
//	MaxNumber: 120,
//}