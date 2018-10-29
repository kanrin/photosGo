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
			signedURL, err := bucket.SignURL(object.Key, oss.HTTPGet, 600, oss.Process("image/auto-orient,1/interlace,1/resize,m_lfit,h_720/quality,q_50"))
			if err != nil {
				HandleError(err)
			}
			r = append(r, Photos{signedURL, 1080, 720})
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