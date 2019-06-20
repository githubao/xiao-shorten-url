// 错误码定义
// author: baoqiang
// time: 2019-06-19 19:54
package main

type ErrNo int64

const OK = 0
const ErrnoParamIllegal = 1
const ErrnoParamUrlFailed = 2
const ErrnoParamNoAvailable = 3

func getErrMsg(code int64) string {
	return ErrMsgMap[code]
}

var ErrMsgMap = map[int64]string{
	OK:                  "成功",
	ErrnoParamIllegal:   "参数错误",
	ErrnoParamUrlFailed: "url格式不对",
	ErrnoParamNoAvailable: "没有可用的短url",
}
