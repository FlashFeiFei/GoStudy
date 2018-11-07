package origin

import "os"



func NewInputDir(data_dir string) inputDir {

	input_dir := inputDir{
		dir:data_dir,
	}
	return input_dir
}

func NewOutDir(data_dir string) inputDir {
	input_dir := inputDir{
		dir: data_dir,
	}
	return input_dir
}

//定义一个操作input目录的结构体
type inputDir struct {
	dir string
}

/*
	判断目录是否存在
 */
func (u inputDir) isDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

/*
	判断目录是否存在，如果存在返回，如果不存在创建目录
 */
func (u inputDir) CreateDir() error {
	//目录名
	is_exists := u.isDirExists(u.dir)
	if is_exists {
		return nil
	} else {
		err := os.MkdirAll(u.dir, os.ModePerm)
		return err
	}
}

/**
	遍历目录下的所有文件
 */
func (u inputDir) ReadDirFile() ([]os.FileInfo, error) {
	file, err := os.Open(u.dir)
	if err != nil {
		return nil, err
	} else {
		file_info, err := file.Readdir(-1)
		return file_info, err
	}
}
