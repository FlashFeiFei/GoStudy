package origin

import (
	"encoding/json"
	"log/domain/parse/result"
	"os"
)

var (
	//文件增量,防止重名
	file_increment int64
)

type OutPutFile struct {
	Item []*result.InputResult `json:"result"`
}

/**
添加一项item
 */
func (of *OutPutFile) AddItem(inputResult *result.InputResult) {
	of.Item = append(of.Item, inputResult)
}

/**
创建输出目录
 */
func (of OutPutFile) createDir(out_dir_root string) (string, error) {
	//目录路径
	dir_string := out_dir_root + "/" + GetYMD()

	dir := NewOutDir(dir_string)
	//创建输出目录
	err := dir.CreateDir()

	if err != nil {
		return "", err
	}
	return dir_string, nil
}

func (of *OutPutFile) write(dir string) error {
	//打开文件
	file_name := GetFileName(&file_increment)
	file_name = dir + "/" + file_name + ".json"
	//创建文件
	file, err := os.Create(file_name)

	if err != nil {
		return err
	}
	//关闭文件
	defer file.Close()

	//json文件写入流
	json_encode := json.NewEncoder(file)
	//用流的方式写入文件
	json_err := json_encode.Encode(of.Item)
	if json_err != nil {
		return json_err
	}
	return nil
}

/*
跑起来
 */
func (of *OutPutFile) run(out_dir_root string) error {
	//创建目录
	dir, err := of.createDir(out_dir_root)
	if err != nil {
		return err
	}
	//写入文件
	write_err := of.write(dir)
	if write_err != nil {
		return write_err
	}

	return nil
}
