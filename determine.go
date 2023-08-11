package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// 判断返回数据包
// i:第几个数据包   bodys:返回包   file:文件名  v:文件后缀  status:返回状态码  body:请求包
func determine(i int, bodys string, file string, v string, status int, body string) {
	//	fmt.Println(status)
	//	fmt.Println(body)
	//判断数据包中是否有文件名
	var characteristic string
	if strings.Contains(bodys, file) {
		color.Red("状态码:%d --  返回包中检测到文件名:%s,上传可能成功 \n", status, file)
		characteristic = characteristic + "返回完整文件名\n"
	}
	//如果包含当前的年月日和文件后缀
	if matchTime(bodys, v) {
		color.Magenta("状态码:%d --  返回包中检测到当前时间+文件后缀,上传可能成功 \n", status)
		characteristic = characteristic + "返回当前时间+文件后缀\n"
	}
	var results []string
	results = append(results, strconv.Itoa(i))
	results = append(results, strconv.Itoa(status))
	results = append(results, body)
	results = append(results, bodys)
	results = append(results, characteristic)
	writer(results)
}
func matchTime(body string, v string) bool {
	currentTime := strings.Replace(time.Now().Format("2006-01-02"), "-", "", -1)
	pattern := currentTime + `.*\.` + v

	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}

	// 进行正则匹配
	if regex.MatchString(body) {
		return true
	} else {
		return false
	}
}
