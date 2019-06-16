package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data)) //编码成字符串，这么函数名有误导性，误以为是编码2字符串
	fmt.Println(sEnc)

	sDec, err := b64.StdEncoding.DecodeString(sEnc)
	if err == nil {
		fmt.Println(string(sDec))
	}

	uEnc := b64.URLEncoding.EncodeToString([]byte(data)) //URL中兼容的base64编码格式
	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

	//这两种编码方式很像，区别在于同一个string编码后base64结束符是+，URL编码是-
}
