package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"github.com/speps/go-hashids"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// salt 盐值
const salt = "CHeF6AC392"

func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64Decode(src string) []byte {
	_bytes, _ := base64.StdEncoding.DecodeString(src)
	return _bytes
}

// Md5String
func Md5String(s string, salt ...string) string {
	h := md5.New()
	if len(salt) > 0 {
		s += strings.Join(salt, "")
	}
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1String
func Sha1String(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256String
func Sha256String(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512String
func Sha512String(s string) string {
	h := sha512.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func HashString(s []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(s, bcrypt.DefaultCost)
	return string(hash)
}

func HashCompare(src, compare []byte) bool {
	return bcrypt.CompareHashAndPassword(src, compare) == nil
}

// HASHIDEncode 混淆
func HASHIDEncode(src int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 10
	h := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{src})
	return e
}

// HASHIDDecode 还原混淆
func HASHIDDecode(src string) int {
	hd := hashids.NewData()
	hd.Salt = salt
	h := hashids.NewWithData(hd)
	e, _ := h.DecodeWithError(src)
	return e[0]
}

// Padding 对明文进行填充
func Padding(plainText []byte, blockSize int) []byte {
	n := blockSize - len(plainText)%blockSize
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

// UnPadding 对密文删除填充
func UnPadding(cipherText []byte) []byte {
	end := cipherText[len(cipherText)-1]
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}

// AESCBCEncrypt AEC加密（CBC模式）
func AESCBCEncrypt(plainText, key, iv []byte) ([]byte, error) {
	//指定加密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//进行填充
	plainText = Padding(plainText, block.BlockSize())
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密连续数据库
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	//返回密文
	return cipherText, nil
}

// AESCBCDecrypt AEC解密（CBC模式）
func AESCBCDecrypt(cipherText, key, iv []byte) ([]byte, error) {
	//指定解密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//解密
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	//删除填充
	plainText = UnPadding(plainText)
	return plainText, nil
}
