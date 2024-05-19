package md5h

import (
	"crypto/md5"
	"encoding/hex"
)

func Enc(input string) string {
	// 创建一个MD5哈希对象
	hasher := md5.New()

	// 将字符串转换为字节数组并计算哈希值
	hasher.Write([]byte(input))

	// 获取哈希值的字节数组
	hashBytes := hasher.Sum(nil)

	// 将哈希值字节数组转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
