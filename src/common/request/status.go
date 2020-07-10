package request

import "github.com/zngue/go_tool/src/fun/zng_str"
//用户删除或者修改状态通用
type IDStatusRequest struct {
	ID int `form:"id" binding:"required"`
	Status int `form:"status"`//1 正常  2 禁用
	From int `form:"form"` // 1 修改状态 2 是删除
	Type int `form:"type"`//来源  0 单个id 1 使用数组 id字符串
	IDArr []int
	IDString string `form:"id_string"`
	Def string `form:"def"` //字符串分隔符 默认 英文逗号分隔
}
func (s *IDStatusRequest ) IDStringToIDArr () *IDStatusRequest  {
	if s.Def!="" {
		s.IDArr = zng_str.IDStringToSlice(s.IDString,s.Def)
	}else{
		s.IDArr = zng_str.IDStringToSlice(s.IDString,",")
	}
	return s
}
func (s *IDStatusRequest)  CheckStatus() *IDStatusRequest  {
	if s.Type==1 {
		s.IDStringToIDArr()
	}else{
		s.IDArr = append(s.IDArr,s.ID)
	}
	return s
}


