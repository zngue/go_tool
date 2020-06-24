package jwt
func CheckLoin(strToken string) error  {
	jwt :=new(JWTAuth)
	_,err :=jwt.Parse(strToken)
	return err
}
