package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GetUUID() string {
	return uuid.NewV4().String()
}

func GetRandomCode(length int) string {
	rand.Seed(time.Now().Unix())

	code := make([]string, 0)

	for i := 0; i < length; i++ {
		code = append(code, fmt.Sprintf("%d", rand.Intn(10)))
	}
	return strings.Join(code, "")
}

func GetRandomString(length int) string {
	bytes := []byte(str)

	result := make([]byte, 0)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
