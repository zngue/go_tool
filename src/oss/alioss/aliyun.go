package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zngue/go_tool/src/config"
	"github.com/zngue/go_tool/src/db"
	"strings"
)
var OssConfig *config.AliyunOss
func init()  {
	OssConfig =&db.Config.AliyunOss
}
func AliOssVersion() string {
	return  oss.Version
}
func AliOssInitServer()(client *oss.Client,err error) {
	client, err = oss.New( OssConfig.Endpoint ,  OssConfig.Accessid ,OssConfig.Accesskey )
	if err != nil {
		return
	}
	return client,err
}

func AliOssUploadFile(localfile string,uploadfile string)(resultfile string,err error){
	resultfile = ""
	// 创建OSSClient实例。
	client,err := AliOssInitServer()

	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	uploadfile = strings.Trim(uploadfile,"/")
	objectName := fmt.Sprintf("%s/%s",OssConfig.Uploaddir,uploadfile) //完整的oss路径
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName := localfile
	// 获取存储空间。
	bucket, err := client.Bucket(OssConfig.Bucket)
	if err != nil {
		return
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return
	}

	if OssConfig.Ssl {
		objectName = "https://"+OssConfig.DomainName+objectName
	}else {
		objectName = "http://"+OssConfig.DomainName+objectName
	}
	resultfile = objectName
	fmt.Println(resultfile)
	return
}
