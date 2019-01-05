package libs

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

type value struct {
	Value int `json:"value,string"`
}

type imageType struct {
	ImageHeight value `json:"ImageHeight"`
	ImageWidth value `json:"ImageWidth"`
}

func ImageInfo(url string) (int, int, error) {
	res, err := http.Get(url)
	HandleErrorF(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	HandleErrorF(err)
	data := &imageType{}
	err = json.Unmarshal(body, data)
	return data.ImageHeight.Value, data.ImageWidth.Value, err
}