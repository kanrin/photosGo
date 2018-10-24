package libs

import (
	"path/filepath"
	"os"
	"path"
	"fmt"
	"encoding/json"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type Items struct {
	Src  string `json:"src"`
	W    int	`json:"w"`
	H    int	`json:"h"`
}

func GetPhotos(host string) []byte {
	f, _ := getFileList(path.Join("photos"))
	return formatJSON(f, host)
}

func getFileList(photoPath string) ([]string, error) {
	var fileList []string
	fileErr := filepath.Walk(photoPath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			} else {
				fileList = append(fileList, path)
			}

			return nil
		})
	return fileList, fileErr
}

func formatJSON (f []string, host string) []byte {
	r := []Items{}
	for _, p := range f {
		file, _ := os.Open(p)
		c, _, _ := image.DecodeConfig(file)
		url := "http://" + host + "/"
		r = append(r, Items{url + p, c.Width, c.Height})
	}
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("%s\n", data)
	}
	return data
}