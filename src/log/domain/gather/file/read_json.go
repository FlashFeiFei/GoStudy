package file

import (
	"encoding/json"
	"os"
)

type ReadJson struct {
	//文件名
	FileName string
	//文件路径
	FileDir string
}

//流的方式读取json数据
func (rj *ReadJson) readStream() error {
	file_path := rj.FileDir + rj.FileName
	//打开文件
	file_json, err := os.Open(file_path)
	if err != nil {
		return err
	}
	//json流
	json_decode := json.NewDecoder(file_json)
}
