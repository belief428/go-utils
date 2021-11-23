package utils

import (
	"regexp"
)

func ValidateMobile(mobile string) bool {
	reg := regexp.MustCompile("^1[3|4|5|6|7|8|9][0-9]\\d{8}$")
	return reg.MatchString(mobile)
}

func ValidateEmail(email string) bool {
	reg := regexp.MustCompile("^[A-Z0-9._%+-]+@[A-Z0-9.-]+\\.[A-Z]{2,6}$")
	return reg.MatchString(email)
}

func ValidateUrl(url string) {
	//reg := regexp.MustCompile("^([hH][tT]{2}[pP]:|||[hH][tT]{2}[pP][sS]:|www\.)(([A-Za-z0-9-~]+)\.)+([A-Za-z0-9-~\/])+$")
}

func ValidateIDCard(obj string) bool {
	reg := regexp.MustCompile("^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$")
	return reg.MatchString(obj)

}

func ValidateIP(ip string) bool {
	reg := regexp.MustCompile("^((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}$")
	return reg.MatchString(ip)
}

func ValidateCompile(obj, compile string) bool {
	reg := regexp.MustCompile(compile)
	return reg.MatchString(obj)
}

func ReplaceAllCompile(obj, compile, replace string) string {
	reg := regexp.MustCompile(compile)
	return reg.ReplaceAllString(obj, replace)
}

func MatchString(pattern string, s string) bool {
	status, _ := regexp.MatchString(pattern, s)
	return status
}
