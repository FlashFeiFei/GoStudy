package parse

import (
	"log"
	"log/domain/parse/origin"
)

func NewParse(out_dir_root string, read_file_dir string) *parse {

	return &parse{
		outDirRoot:out_dir_root,
		readFileDir:read_file_dir,
	}
}

type parse struct {
	//输出根目录
	outDirRoot string
	//读取的文件名的路径
	readFileDir string
}

//实现线程池的任务接口
func (p *parse) Task() {
	input_file := origin.NewInputFile(p.readFileDir)
	//执行解析文件
	err := input_file.ReadFile(p.outDirRoot)

	if err != nil {
		log.Println("解析文件警告:", err)
	}
}
