package main

import (
	"log"
	readjson "log/domain/gather/file"
	"log/domain/parse/origin"
	"log/domain/parse/result"
	"sync"
)

func main() {

	input_dir_path := "data/output"

	input_dir := origin.NewInputDir(input_dir_path)

	file_info_list, err := input_dir.ReadDirFile()
	if err != nil {
		//目录打开失败
		log.Fatalln(err)
	}

	for _, item := range file_info_list {
		if item.IsDir() {
			//遍历目录下的所有文件
			file_dir_path := input_dir_path + "/" + item.Name()
			file_dir := origin.NewInputDir(file_dir_path)
			file_list, err := file_dir.ReadDirFile()
			if err != nil {
				continue
			}

			file_count := len(file_list)

			result_chan := make(chan *result.InputResult)
			over := make(chan int)

			//读取文件，进入通道
			for _, file := range file_list {

				go func(input_file_name string) {
					read_json := readjson.ReadJson{
						FileName:   input_file_name,
						FileDir:    file_dir_path,
						ResultChan: result_chan,
						OverChan:   over,
					}
					read_json.Task()
				}(file.Name())

			}

			//判断通道发送数据是否完毕，完毕的话，关闭通道
			go func(count_total int) {
				count := 0
				for  {
					select {
					case <-over:
						count++
						if count_total == count {
							close(result_chan)
							close(over)
						}
					default:
						continue
					}
				}
			}(file_count)

			var wg sync.WaitGroup
			//一个录入数据线程,一个写
			wg.Add(1)
			////将通道中的文件全部汇总起来
			go func() {
				write_json := readjson.WriteJson{
					NameDate:   item.Name(),
					ResultChan: result_chan,
				}
				//写入数据
				write_json.Write()
				wg.Done()
			}()
			wg.Wait()
			//遍历下一个目录下的
		}
	}
	log.Println("程序结束了")
}
