package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func converToBianry(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

// 1,2,4
func TestExchange(t *testing.T) {
	for i := 0; i < 10; i++ {
		//t.Log(1 << uint(i))
	}

	//t.Log(0b00000000100)
	a := 1
	b := 1 << 1
	c := 1 << 2
	t.Log(a)
	t.Log(b)
	t.Log(c)

	d := a | b | c
	t.Log(d)

	t.Log(d & a)
	t.Log(d & b)
	t.Log(d & c)
}

func TestAnyToByte(t *testing.T) {
	src := []int{1}

	mark := 0

	for i := 0; i < len(src); i++ {
		mark = mark | src[i]
	}
	t.Log(mark)
}

func TestFromJSONFile(t *testing.T) {
	file := "FWBAT-GX-A13-V310.bin"
	f, err := os.Open(file)
	if err != nil {
		t.Log(err)
		return
	}
	defer f.Close()
	content, err := ioutil.ReadAll(f)

	if err != nil {
		t.Log(err)
		return
	}
	t.Log(content)
	t.Log(len(content))

	inputReader := bufio.NewReader(f)
	s, _, _ := inputReader.ReadLine()

	t.Log(s)
	t.Log(len(s))
	//t.Log(bufio.NewReader(f))
	//for {
	//	inputString, readerError := inputReader.ReadString('\n') //我们将inputReader里面的字符串按行进行读取。
	//	if readerError == io.EOF {
	//		return
	//	}
	//	t.Log(inputString)
	//}
}
