package utils

import (
	"testing"
)

func TestReplaceCompile(t *testing.T) {
	src := "12312321321"

	t.Log(src)

	compile := "[/]+"
	src = ReplaceAllCompile(src, compile, "/")
	t.Log(src)

	compile = "(^[\\w])([\\w/]*)([\\w])$"
	//compile = "(^[\\w])"

	//compile = "^[\\w]" +
	//	"([\\w*])" +
	//	"[\\w]$"
	status := ValidateCompile(src, compile)
	t.Log(status)

	//compile := "user.+\\z"
	//src = ReplaceAllCompile(src, compile, "user")
	//t.Log(src)

	//compile = "^\\w{0,50}$"
	//status = ValidateCompile(src, compile)
	//t.Log(status)
	//
	////支持中文、大小写字母、日文、数字、短划线、下划线、斜杠和小数点，必须以中文、英文或数字开头，不超过 30 个字符
	//compile = "(^[a-zA-Z0-9\u4e00-\u9fa5])([a-zA-Z0-9_\u4e00-\u9fa5-_/.]*){0,30}$"
	//status = ValidateCompile(src, compile)
	//t.Logf("文本检测：%v\n", status)
}

func TestValidateCompile(t *testing.T) {
	src := "2134"

	t.Log(src)

	compile := "^[\\w-.:]{4,32}$"
	status := ValidateCompile(src, compile)
	t.Log(status)
}
