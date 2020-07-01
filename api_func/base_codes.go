package api_func

const (
	OK       = 2000 //请求成功
	ArgError = 3001 //请求参数错误
)

var statusText = map[int]string{
	OK:       "OK",
	ArgError: "Arg Error",
}

//resulf message
func Msg(code int) string {
	return statusText[code]
}
