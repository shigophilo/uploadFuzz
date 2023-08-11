package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// 发送请求
func send(url string, list map[string]string, bd string, fileSuffix []string) {
	body := bd
	for _, v := range fileSuffix {
		rangs := randCreator(Len)
		file := rangs + "." + v
		body = strings.Replace(body, `#filename#`, file, -1)
		var req *http.Request
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		req, _ = http.NewRequest("POST", url, strings.NewReader(body))

		for k, vv := range list {
			if k == "" || vv == "" {
				continue
			}
			//	fmt.Println(k, vv)
			req.Header.Add(k, vv)
		}
		//fmt.Println(i)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(url + "-- request fail")
			return
		}
		defer resp.Body.Close()
		bodys, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取页面返回失败")
		} else {
			file := rangs + "." + v
			fmt.Printf("请求包:%d -- 状态码:%d -- 文件名:%s \n", Nor, resp.StatusCode, file)
			Nor++
			determine(Nor, string(bodys), file, v, resp.StatusCode, body)
		}
		//数据包中的#filename#已经被替换过了,先恢复成原来的数据包#filename#,
		body = bd
	}
}

func randCreator(l int) string {
	str := "0123456789abcdgklmnopqrstuvwxyz"
	strList := []byte(str)

	result := []byte{}
	i := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i = i + 1
	}
	return string(result)
}
