package libs

import (
	"github.com/kylelemons/go-gypsy/yaml"
)

func GetOssInfo() (string, string, string, string){
	config, err := yaml.ReadFile("config.yml")
	HandleErrorF(err)
	endPoint, err := config.Get("endPoint")
	HandleErrorF(err)
	accessKeyId, err := config.Get("accessKeyId")
	HandleErrorF(err)
	accessKeySecret, err := config.Get("accessKeySecret")
	HandleErrorF(err)
	bucketName, err := config.Get("bucketName")
	HandleErrorF(err)
	return endPoint, accessKeyId, accessKeySecret, bucketName
}