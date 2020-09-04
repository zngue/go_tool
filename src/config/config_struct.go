package config

import "crypto/rsa"

type Config struct {
	Mysql Mysql `json:"mysql" `
	Redis Redis	`json:"redis.yaml" `
	System System `json:"system" `
	HttpRequest HttpRequest  `json:"http_request" `
	ServiceList []ServiceList `json:"serviceList" `
	DefaultLoad DefaultLoad `json:"defaultLoad"`
	AliyunOss AliyunOss `json:"aliyunoss" `
	JWT Jwt `json:"jwt"`
	AliAppClient AliAppClient `json:"ali_app_client"`
	WeChat WeChat 	`json:"wechat"`
}
//数据库配置信息
type Mysql struct {
	DBName string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`
	Config string	`json:"config"`
	Host   string `json:"host"`
	Port	string `json:"port"`
	LogMode bool `json:"logMode"`
	MaxIdleConns int `json:"maxIdleConns"`
	MaxOpenConns int `json:"maxOpenConns"`
	Charset string `json:"charset"`
	TimeStamp bool `json:"timeStamp"`
}
type AliAppClient struct {
	PartnerID  string  `json:"partner_id"`//合作者ID
	SellerID   string `json:"seller_id"`
	AppID      string `json:"app_id"` // 应用ID
	PrivateKey *rsa.PrivateKey `json:"private_key"`
	PublicKey  *rsa.PublicKey `json:"public_key"`
}
//redis配置信息
type Redis struct {
	Host   string `json:"host"`
	Port	string `json:"port"`
	Password string `json:"password"`
	DBNum	int `json:"dbNum"`
	PoolSize int `json:"poolSize"`
}
//系统配置信息
type System struct {
	Port string `json:"port"`
}
//http请求配置信息
type HttpRequest struct {
	ServiceMode	bool `json:"serviceMode"`
	RedisDBNum int `json:"redisDbNum"`
}
//http 服务请求配置
type ServiceList struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
//默认加载数据
type DefaultLoad struct {
	Mysql bool `json:"mysql"`
	Redis bool `json:"redis.yaml"`
}
type AliyunOss struct {
	Accessid string `json:"accessid"`
	Accesskey string `json:"accesskey"`
	Endpoint string `json:"endpoint"`
	Bucket string `json:"bucket"`
	Uploaddir string `json:"uploaddir"`
	DomainName string `json:"domainName"`
	Ssl bool `json:"ssl"`
}
type Jwt struct {
	Secret string `json:"secret"`
	ExpireTime int `json:"expireTime"`
	Issuer string `json:"issuer"`
	Subject string `json:"subject"`
}
type WeChat struct {
	AppID string `json:"appID"`
	AppSecret string `json:"appSecret"`
	Token string `json:"token"`
	EncodingAESKey string `json:"encodingAESKey"`
	PayMchID string `json:"payMchID"`
	PayNotifyURL string `json:"payNotifyURL"`
	PayKey string `json:"payKey"`
}

