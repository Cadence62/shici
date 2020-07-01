package api_func

// 返回值结构定义
type Result struct {
	Code   int         `json:"code"`
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Total  int         `json:"total"`
	Data   interface{} `json:"data"`
}

func NewResult() *Result {
	return &Result{
		Code: 0,
		Data: nil,
	}
}
