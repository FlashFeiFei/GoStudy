package main

import (
	"log"
	"log/domain/parse"
	"log/domain/parse/origin"
	"log/domain/parse/work"
	"sync"
)

func main() {
	//数据存放路径
	input_dir_path := "data/input"

	input_dir := origin.NewInputDir(input_dir_path)
	//创建目录
	input_dir.CreateDir()

	//遍历存放数据的目录下所有的文件
	file_info_list, err := input_dir.ReadDirFile()
	if err != nil {
		log.Fatalln(err)
		return
	}

	//统计有目录下有多少个文件(目录不算)
	work_total := 0
	for _, file_info := range file_info_list {
		is_dir := file_info.IsDir()
		if is_dir {
			//目录的时候跳过
			continue
		}
		work_total++
	}

	//开启一个5个线程的线程池
	pool_work := work.New(5)
	var wg sync.WaitGroup
	wg.Add(work_total)
	//输出文件的根目录
	out_dir_root := "data/output"
	//一个文件用一个线程去执行解析
	//因为用了线程池所以同一时最多有5个线程同时运行
	//剩下的必须等待前面的任务完成
	for _, file_info := range file_info_list {
		is_dir := file_info.IsDir()
		if is_dir {
			//目录的时候跳过
			continue
		}

		//读取文件的路径
		read_file_dir := input_dir_path + "/" + file_info.Name()

		go func() {
			pase_work := parse.NewParse(out_dir_root, read_file_dir)
			//进入任务进入线程池
			//重点！！Run方法必须是指针才行
			pool_work.Run(pase_work)
			wg.Done()
		}()
	}

	//等待所有任务进入线程池
	wg.Wait()
	//关闭通道
	//关闭通道并等待通道中的所有任务运行完毕
	pool_work.Shutdown()
}
