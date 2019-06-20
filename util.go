// 工具函数
// author: baoqiang
// time: 2019-06-19 20:01
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func Md5HexFromString(data, salt string) []byte {
	hash := md5.New()
	hash.Write([]byte(data))
	if len(salt) > 0 {
		hash.Write([]byte(salt))
	}
	return []byte(hex.EncodeToString(hash.Sum(nil)))
}

func URLShorten(longURL string) []string {
	key := "URLShorten"
	text := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	md5text := Md5HexFromString(longURL, key)

	shortURLs := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		str := md5text[i*8 : (i+1)*8]

		// str to int
		num, err := strconv.Atoi(fmt.Sprintf("%x", string(str)))
		if err != nil {
			panic(fmt.Errorf("invalid longurl: %v", longURL))
		}

		// 选择低30位
		num &= 0x3FFFFFFF

		// 取30位的后6位与0x0000003D进行逻辑与操作，结果范围是0~61，作为text的下标选择字符
		// 把num左移5位重复进行，得到6个字符组成短URL
		shortURL := make([]byte, 0, 6)
		for j := 0; j < 6; j++ {
			shortURL = append(shortURL, text[num&0x0000003D])
			num >>= 5
		}

		// add one
		shortURLs = append(shortURLs, string(shortURL))
	}

	return shortURLs
}
