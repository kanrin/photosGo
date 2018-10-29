package libs

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

func GetOssInfo() (string, string, string, string){
	config, _ := yaml.ReadFile("config.yml")
	endPoint, _ := config.Get("endPoint")
	accessKeyId, _ := config.Get("accessKeyId")
	accessKeySecret, _ := config.Get("accessKeySecret")
	bucketName, _ := config.Get("bucketName")
	return endPoint, accessKeyId, accessKeySecret, bucketName
}