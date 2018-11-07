package result

type InputResult struct {
	//重哪里来
	FromIp string `json:"fromIp"`
	Domain string `json:"domain"`
	//图片大小
	Size string `json:"size"`
}

func (input *InputResult) Empty() bool {
	if input.Domain == "" {
		return true
	}
	if input.FromIp == "" {
		return true
	}
	if input.Size == "" {
		return true
	}
	return false
}
