package base64

import (
	"encoding/base64"
)

// Base64Encode base64 standard decoding (标准方式编码)
func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// Base64Decode base64 standard encoding (标准方式解码)
func Base64Decode(input string) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Base64UrlEncode base64 url decoding (安全方式编码)
// URL和文件名安全方式是标准方式的变体，其输出用于URL和文件名。
// 因为+和/字符是标准Base64字符对URL和文件名编码不安全，变体即使用-代替+，_（下划线）代替/ 。
func Base64UrlEncode(input []byte) string {
	return base64.URLEncoding.EncodeToString(input)
}

// Base64UrlDecode base64 url encoding (安全方式解码)
func Base64UrlDecode(input string) ([]byte, error) {
	result, err := base64.URLEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}