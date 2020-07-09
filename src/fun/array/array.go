package array

func Uniq()  {
	
}

//判断是否int是否在一个数组中
func IntIsArray(int int,intArr []int) bool  {
	if len(intArr)==0 {
		return false
	}
	for _,val :=range intArr{
		if  val== int{
			return true
		}
	}
	return false
}

func  StringIsInArray( str string ,strArr  []string ) bool {
	if len(strArr)==0 {
		return false
	}
	for _,val :=range strArr{
		if val!=" " && str!=" " && val== str{
			return true
		}
	}
	return false
}
