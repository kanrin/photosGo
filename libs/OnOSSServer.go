package libs

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"encoding/json"
)

type Photos struct {
	Src  string `json:"src"`
	W    int	`json:"w"`
	H    int	`json:"h"`
}

func GetPhotosOss() []byte {
	e, ai, as, bn := GetOssInfo()
	client, err := oss.New(e, ai, as)
	if err != nil {
		HandleError(err)
	}
	bucket, err := client.Bucket(bn)
	if err != nil {
		HandleError(err)
	}

	// 列举所有文件。
	marker := ""
	r := []Photos{}
	for {
		lsRes, err := bucket.ListObjects(oss.Marker(marker))
		if err != nil {
			HandleError(err)
		}

		for _, object := range lsRes.Objects {
			signedURL := "http://localhost:8000/get/" + object.Key
			r = append(r, Photos{signedURL, 300, 200})
		}

		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	data, err := json.Marshal(r)
	if err != nil {
		HandleError(err)
	}
	return data
}