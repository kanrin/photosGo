package libs

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
)

func getOssInfo() (string, string, string){
	config, _ := yaml.ReadFile("config.yml")
	endPoint, _ := config.Get("endPoint")
	accessKeyId, _ := config.Get("accessKeyId")
	accessKeySecret, _ := config.Get("accessKeySecret")
	return endPoint, accessKeyId, accessKeySecret
}

func GetPhotosOss() {
	e, ai, as := getOssInfo()
	client, err := oss.New(e, ai, as)
	if err != nil {
		fmt.Errorf("[OSSError]: %s", err)
	}
	lsRes, err := client.ListBuckets()
	if err != nil {
		fmt.Errorf("[OSSError]: %s", err)
	}
	fmt.Print(e)
	fmt.Println(lsRes)
	for _, bucket := range lsRes.Buckets {
		fmt.Println(bucket.Name)
	}
}