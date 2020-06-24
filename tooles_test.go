package main

import (
	"encoding/json"
	"github.com/zngue/go_tool/src/idea"
	"github.com/zngue/go_tool/src/jwt"
	"testing"
	"time"
)

func TestLog(t *testing.T)  {


}

func TestIdea(t *testing.T)  {

	idea.GetAllIdea()
}
func TestJwt( t *testing.T)  {

	jwt :=jwt.JWTAuth{}

	//token ,_:=jwt.CreateToken("list")
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vdGVzdGFwaS5jcWNiLmNvbS9hcGkvdXNlci9jb2RlX2xvZ2luIiwiaWF0IjoxNTg5ODg3NjQ4LCJleHAiOjE5NDk4ODc2NDgsIm5iZiI6MTU4OTg4NzY0OCwianRpIjoiWDlWakttWWFBQkx6UkFuTSIsInN1YiI6ODc0MTk4LCJwcnYiOiI4NjY1YWU5Nzc1Y2YyNmY2YjhlNDk2Zjg2ZmE1MzZkNjhkZDcxODE4In0.mzzlph4FWpS3f1XtvABLPg15sByKTw-1o9i_STdBOSI"
	ss ,err:=jwt.Parse(token)
	t.Log(err)
	time.Sleep(11)
	t.Log(ss)

}

type User struct {
	ID int
	Name string
	
}
func teName(i interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	b,_:=json.Marshal(i)
	json.Unmarshal(b,&m)
	return m
}
func TestUserList(t *testing.T) {
	user :=User{

		10,
		"zhansgan",
	}

	t.Log(teName(user))
}