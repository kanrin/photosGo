package libs

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
)

func GetOssFile(name string) ([]byte, error) {
	e, ai, as, bn := GetOssInfo()
	client, err := oss.New(e, ai, as)
	if err != nil {
		HandleError(err)
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		HandleError(err)
	}
	body, err := bucket.GetObject(name)
	if err != nil {
		HandleError(err)
	}
	defer body.Close()
	return ioutil.ReadAll(body)
}