package ut

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	res := Md5("mc")
	fmt.Println(res, "=======")
}

func TestBase64(t *testing.T) {
	res := Base64Enc("mcdj")
	fmt.Println(res)
	res1, err := Base64Dec(res)
	fmt.Println(res1, err, "====------")
}