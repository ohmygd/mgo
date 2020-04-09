package ut

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"
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

func IsMobile(mobile string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)

	return rgx.MatchString(mobile)
}

func RandStr(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var res string
	for i:=0;i<n;i++ {
		rand.Seed(time.Now().UnixNano())
		res += string(str[rand.Intn(len(str))])
	}

	return res
}

// 密码长度需不小于8位
func PwdIsOk(pwd string) bool {
	if len(pwd) < 8 {
		return false
	}

	return true
}

func Hello() {
	fmt.Println("hello--------")
}