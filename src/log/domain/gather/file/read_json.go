package file

import (
	"encoding/json"
	"log"
	"log/domain/parse/result"
	"os"
)

type ReadJson struct {
	//文件名
	FileName string
	//文件路径
	FileDir string

	ResultChan chan *result.InputResult

	OverChan chan int
}

//流的方式读取json数据
func (rj *ReadJson) Task() {
	file_path := rj.FileDir + "/" + rj.FileName
	//打开文件
	file_json, err := os.Open(file_path)
	if err != nil {
		return
	}
	//读取json
	json_decode := json.NewDecoder(file_json)

	var input_result []*result.InputResult
	json_decode.Decode(&input_result)
	//写入通道
	for _, item := range input_result {
		rj.ResultChan <- item
	}
	//发送一个结束信号
	rj.OverChan <- 1

	log.Println("输入结束")
	return
}
