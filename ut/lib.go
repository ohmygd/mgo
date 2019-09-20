package ut

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"time"
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

func StrComplex(str string) string {
	prefix1 := strconv.Itoa(rand.Intn(1000))
	prefix2 := strconv.Itoa(rand.Intn(1000))

	return prefix1 + prefix2 + str + strconv.FormatInt(time.Now().Unix(), 10)
}