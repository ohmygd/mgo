package ut

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)

	return md5str1
}

func Base64Enc(str string) string {
	encodeString := base64.StdEncoding.EncodeToString([]byte(str))

	return encodeString
}

func Base64Dec(str string) (res string, err error) {
	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}

	res = string(decodeBytes)
	return
}
