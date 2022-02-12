package tool

//CanUse 判断是否为合法数据
//参数 输入的数据
//操作 不允许含有特殊符号
//返回值 是否允许登录的bool值
func CanUse(str string) bool {
	for _,v := range []byte(str){
		if v=='#'||v=='$'||v=='%'||v=='^'||v=='&'||v=='*'||v=='='{
			return false
		}
	}
	return true
}