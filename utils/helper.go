package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// pkcs7Padding 填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// AesEncrypt 加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//判断加密快的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	encryption := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(encryption, encryptBytes)
	return encryption, nil
}

// AesDecrypt 解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	encryption := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(encryption, data)
	//去除填充
	encryption, err = pkcs7UnPadding(encryption)
	if err != nil {
		return nil, err
	}
	return encryption, nil
}

// EncryptByAes Aes加密 后 base64 再加
func EncryptByAes(data string, pwdKey string) (string, error) {
	res, err := AesEncrypt([]byte(data), []byte(pwdKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(res), nil
}

// DecryptByAes Aes 解密
func DecryptByAes(data string, pwdKey string) (string, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decryption, err := AesDecrypt(dataByte, []byte(pwdKey))
	if err != nil {
		return "", err
	}
	return string(decryption), nil
}

func GenerateRandomString() string {
	// 定义包含的字符集合
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-=_+[]{}|;:'\",.<>/?`~"

	// 生成随机字符串
	var result strings.Builder
	for i := 0; i < 32; i++ {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		result.WriteByte(charSet[randomIndex.Int64()])
	}

	return result.String()
}

// TokenMd5 md5获取token
func TokenMd5() string {
	curTime := time.Now().Unix()
	//fmt.Println("curTime", curTime)
	h := md5.New()
	//fmt.Println("h-->", h)
	//fmt.Println("strconv.FormatInt(curTime, 10)-->", strconv.FormatInt(curTime, 10))
	io.WriteString(h, strconv.FormatInt(curTime, 10))
	//fmt.Println("h-->", h)
	token := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("token--->", token)
	return token
}
