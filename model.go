// model
// author: baoqiang
// time: 2019-06-20 11:36
package main

type Response struct {
	ErrCode int64  `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Result  string `json:"result"`
}
