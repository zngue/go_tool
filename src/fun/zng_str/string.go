package zng_str

import (
	"strconv"
	"strings"
)

//@param idsString id字符串
//@param def 分隔符
//@return idsArr id切片
func IDStringToSlice(idsString string,def string) ( idsArr []int)  {
	ids:=strings.Split(idsString,def)
	for _,v:=range ids{
		id,err:=strconv.Atoi(v)
		if err ==nil {
			idsArr=append(idsArr, id)
		}
	}
	return
}
//@param str 字符
//@return int id int类型
func StringToInt( str string ) (int int,err error)  {
	int,err=strconv.Atoi(str)
	return
}
//@param int2 数字
func IntToString( int2 int ) string {
	return  strconv.Itoa(int2)
}

