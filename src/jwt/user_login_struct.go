package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/zngue/go_tool/src/db"
	"time"
)

type UserInfo struct {
	DataInfo interface{}
}

type JWTClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserInfo UserInfo `json:"user_info"`
}
var (
	Secret   string   // 加盐
	ExpireTime int        // token有效期
	Issuer string
	Subject string
)
const(
	ErrorreasonServerbusy = "服务器繁忙"
	ErrorreasonLoginOutTime= "登录过期，请重新登录"
)

func init()  {
	jwtConfig  := db.Config.JWT
	Secret = jwtConfig.Secret
	ExpireTime =jwtConfig.ExpireTime
	Issuer =jwtConfig.Issuer
	Subject = jwtConfig.Subject
}
type JWTAuth struct {

}
func (JWTAuth) CreateClaims (data  interface{}) (claims *JWTClaims) {

	claims = &JWTClaims{
		UserInfo: UserInfo{
			DataInfo: data,
		},
		StandardClaims:jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix(),
			Issuer: Issuer,
			Subject:Subject,
		},
	}
	return
}
func (j *JWTAuth) CreateToken( data  interface{})(token string,err error)  {
	claims:=j.CreateClaims(data)
	token,err=j.GetToken(claims)
	return
}
func (JWTAuth) GetToken(claims *JWTClaims)(signedToken string,err error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(Secret))
	if err != nil {
		err=errors.New(ErrorreasonServerbusy)
	}
	return
}
func (JWTAuth) Parse(strToken string) (claims *JWTClaims,err error) {
	token, errs := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if errs != nil {
		return nil,errs
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		err= errors.New(ErrorreasonLoginOutTime)
		return nil,err
	}
	if err := token.Claims.Valid(); err != nil {
		return nil,err
	}
	return claims,nil
}
func (j *JWTAuth) Refresh(token string) (newToken string, err error){
	claims,err := j.Parse(token)
	if err!=nil {
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	newToken,err=j.GetToken(claims)
	if err!=nil {
		return
	}
	return
}


