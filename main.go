package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/qax-os/excelize"
)

var Outfile string

// Number of requests
var Nor int
var Len int

func main() {
	var language string
	var input string
	var url string
	flag.StringVar(&language, "l", "all", "脚本语言 a p j all")
	flag.StringVar(&input, "i", "input.txt", "文件上传数据包")
	flag.StringVar(&url, "u", "", "url")
	flag.StringVar(&Outfile, "o", "outfile", "结果输出文件名")
	flag.IntVar(&Len, "len", 9, "文件名长度")
	flag.Parse()
	start()
	if url == "" {
		fmt.Println("请设置文件上传url地址  -u")
		os.Exit(0)
	}
	//获取请求头
	list := analysis(input)
	//获取请求体
	body := getBody(input)
	//获取后缀
	fileSuffix := Suffix(language)
	//多后缀
	fmt.Println("开始多后缀")
	send(url, list, body, fileSuffix)
	//截断字符
	filetruncation := truncation(fileSuffix)
	fmt.Println("开始截断")
	send(url, list, body, filetruncation)
	//后缀名大小写
	fmt.Println("开始大小写")
	fileCase := suffixCase(fileSuffix)
	send(url, list, body, fileCase)
	//文件名后缀双写
	fmt.Println("文件名后缀双写")
	filedouble := doubleWriting(fileSuffix)
	send(url, list, body, filedouble)
	//修改Content-Type
	fmt.Println("修改Content-Type")
	bodys := modifyContent(body)
	for _, v := range bodys {
		send(url, list, v, fileSuffix)
	}
	//	fmt.Println(body)
	deleteBlankLines(Outfile)
}

// 获取请求体
func getBody(input string) string {
	//defer file.Close()
	file, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("数据包文件读取失败(getbody),请使用 -i", err)
		os.Exit(0)
	}
	body := string(file)
	start := strings.Index(body, "\r\n\r\n")
	end := len(body)
	body = body[start+4 : end]
	if !strings.Contains(body, "#filename#") {
		red := color.New(color.FgHiWhite)
		whiteBackground := red.Add(color.BgHiMagenta)
		whiteBackground.Println("请求包中未发现文件名变量,将发送默认数据包,如测试多后缀,请将上传文件名修改成\"#filename#\"")
	}
	return body
}

// 解析请求头
func analysis(input string) map[string]string {
	request_header := make(map[string]string)
	var arr []string
	file, err := os.Open(input)
	if err != nil {
		fmt.Println("数据包文件读取失败,请使用 -i")
		os.Exit(0)
	}
	defer file.Close()
	readfile := bufio.NewReader(file)
	for {
		line, err := readfile.ReadString('\n')
		arr = strings.Split(line, ":")
		var key string
		var value string
		for i, v := range arr {
			if strings.Contains(strings.ToLower(v), "content-Length") {
				continue
			}
			if i == 0 {
				key = strings.TrimSpace(v)
			} else {
				value = strings.TrimSpace(v)
			}
		}
		request_header[key] = value
		if line == "\r\n" || len(line) == 0 || line == "" || err == io.EOF {
			break
		}
	}
	//	fmt.Println(request_header)
	return request_header
}

func deleteBlankLines(outfile string) {
	file := Outfile + ".xlsx"
	f, err := excelize.OpenFile(file)
	if err != nil {
		fmt.Println(err)
	}
	err = f.RemoveRow("上传详情", 2)
	if err != nil {
		fmt.Println(err)
	}
	if err := f.SaveAs(file); err != nil {
		fmt.Println(err)
	}
}

func start() {
	fmt.Println("=====================================================================================")
	color.Cyan("              _                                     _                     _       ")
	color.Red("             | |                             _     ( )     _             | |      ")
	color.Yellow("  ____   ____| |__   ___  ____  _____  ___ _| |_   |/    _| |_ ___   ___ | |  ___ ")
	color.Blue(" |    \\ / ___)  _ \\ / _ \\|  _ \\| ___ |/___|_   _)       (_   _) _ \\ / _ \\| | /___)")
	color.Magenta(" | | | | |   | | | | |_| | | | | ____|___ | | |_          | || |_| | |_| | ||___ |")
	color.Green(" |_|_|_|_|   |_| |_|\\___/|_| |_|_____|___/   \\__)          \\__)___/ \\___/ \\_|___/ ")
	fmt.Println("=====================================================================================")
	red := color.New(color.FgRed)
	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("请删除文件上传数据包的第一行内容")
}
