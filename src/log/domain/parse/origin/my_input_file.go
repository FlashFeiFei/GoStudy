package origin

import (
	"bufio"
	"io"
	"log"
	"log/domain/parse/match"
	"log/domain/parse/result"
	"os"
)

func NewInputFile(file_name string) inputFile {
	input_file := inputFile{
		name: file_name,
	}
	return input_file
}

type inputFile struct {
	//文件的路径
	name string
}

/*
从文件中读取一条数据，然后写入文件
 */
func (fi inputFile) ReadFile(out_dir_root string) (error) {
	//打开文件
	file, err := os.Open(fi.name)
	if err != nil {
		return err
	}
	//关闭文件
	defer file.Close()
	//以流的方式读取
	rd := bufio.NewReader(file)
	var out_put_file *OutPutFile
	//输出源
	out_put_file = &OutPutFile{}
	for {
		//从中解析出的一条数据
		var input_result *result.InputResult
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			//读取遇到失败或者已经读取完数据
			//停止读取
			break
		}
		input_match := match.NewInputMatch(line)
		input_result, err = input_match.Match()
		if err != nil {
			//发生错误，跳过这条数据
			continue
		}
		//添加到输出源
		out_put_file.AddItem(input_result)
	}

	//写到输出源
	run_err := out_put_file.run(out_dir_root)

	if run_err != nil {
		//写入出错
		log.Println(run_err)
		return run_err
	}
	return nil
}
