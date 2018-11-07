package match

import (
	"errors"
	"log/domain/parse/result"
	"strings"
)

func NewInputMatch(data string) inputMatch {
	return inputMatch{
		line_data: data,
	}
}

type inputMatch struct {
	//一行数据
	line_data string
}

//设置匹配的数据
func (match *inputMatch) SetData(data string) {
	match.line_data = data
}

func (match *inputMatch) Match() (*result.InputResult, error) {
	//按照空格切割字符串
	var input_result *result.InputResult
	//构造一个inputResult结构体
	input_result = &result.InputResult{}
	data := strings.Fields(match.line_data)
	for index, item := range data {
		switch index {
		case 0:
			//ip位置
			input_result.FromIp = item
		case 9:
			//图片大小
			input_result.Size = item
		case 10:
			//通过请求路径获取域名
			domian, err := match.domianParse(item)
			if err != nil {
				//停止匹配
				continue
			}
			input_result.Domain = domian
		default:
			//跳过这项数据
			continue
		}
	}
	//判断结构体是否为空
	is_empty := input_result.Empty()
	if !is_empty {
		//不为空的时候
		return input_result, nil
	}

	return nil, errors.New("空域名异常")
}

func (match inputMatch) domianParse(domian string) (string, error) {
	is_domian := strings.Contains(domian, "http")
	//域名的时候才开始匹配
	if is_domian {
		var result string
		domian_split := strings.Split(domian, "//")
		domin_string := strings.Split(domian_split[1], "/")
		result = domin_string[0]
		return result, nil
	}
	//传进来的不是域名,直接返回
	return "", nil
}
