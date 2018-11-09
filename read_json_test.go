package main

import (
	"log/domain/gather/file"
	"log/domain/parse/result"
	"testing"
)

func TestReadJson(t *testing.T){

	result_chan := make(chan *result.InputResult)

	read_json := file.ReadJson{
		FileDir:"data/output/2018-11-09/",
		FileName:"1541738116580489700_1.json",
		ResultChan:result_chan ,
	}
	read_json.ReadStream()
}
