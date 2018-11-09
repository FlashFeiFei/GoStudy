package file

import (
	"encoding/json"
	"log"
	"log/domain/parse/result"
	"os"
)

type WriteJson struct {
	NameDate   string
	ResultChan chan *result.InputResult
}

func (wj *WriteJson) Write() {
	file_dir := "data/gather/" + wj.NameDate
	dir_exists := DirExists(file_dir)
	if !dir_exists {
		//创建目录
		os.MkdirAll(file_dir, os.ModePerm)
	}

	file_name := file_dir + "/gather.json"
	file_exists := FileExists(file_name)
	if file_exists {
		//文件存在的时候，删除文件
		err := os.Remove(file_name)
		if err != nil {
			log.Fatalln(err)
		}
	}
	//创建文件
	file, err := os.Create(file_name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var result_gather []*result.InputResult
	json_encode := json.NewEncoder(file)
	for {

		item, ok := <-wj.ResultChan
		//log.Println(item)
		if !ok {
			log.Println("通道没有数据")
			break
		}
		result_gather = append(result_gather, item)
	}
	log.Println("写入文件")
	json_encode.Encode(&result_gather)
}
