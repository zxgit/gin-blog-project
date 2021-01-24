package errcode

var (
	Success                   = NewErrow(0, "ok")
	ServerError               = NewErrow(10000000, "服务器内部错误")
	InvalidParams             = NewErrow(10000001, "入参错误")
	NotFound                  = NewErrow(10000002, "数据不存在")
	UnauthorizedAuthNotExist  = NewErrow(10000003, "鉴权失败，找不到对应的appKey和appSecret")
	UnauthorizedTokenError    = NewErrow(10000004, "鉴权失败，token错误")
	UnauthorizedTokenTimeout  = NewErrow(10000005, "鉴权失败，token超时")
	UnauthorizedTokenGenerate = NewErrow(10000006, "鉴权失败，token生成失败")
	TooManyRequests           = NewErrow(10000007, "q请求过多")
)
