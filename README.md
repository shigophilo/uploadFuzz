- 原始数据包

```
POST /saveupload.asp HTTP/1.1
Host: 127.0.0.1
Content-Length: 215
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
Origin: null
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7yyQ5XLHOn6WZ6MT
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
Connection: close

------WebKitFormBoundary7yyQ5XLHOn6WZ6MT
Content-Disposition: form-data; name="fname"; filename="1.asp"
Content-Type: image/jpeg

aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa?
------WebKitFormBoundary7yyQ5XLHOn6WZ6MT--

```

- 修改

```
//此行删除 POST /saveupload.asp HTTP/1.1
Host: 127.0.0.1
Content-Length: 215
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
Origin: null
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7yyQ5XLHOn6WZ6MT
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
Connection: close

------WebKitFormBoundary7yyQ5XLHOn6WZ6MT
Content-Disposition: form-data; name="fname"; filename="#filename#"  //替换文件名变量
Content-Type: image/jpeg

aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa?
------WebKitFormBoundary7yyQ5XLHOn6WZ6MT--

```

- 修改后  input.txt

```
Host: 127.0.0.1
Content-Length: 215
Cache-Control: max-age=0
Upgrade-Insecure-Requests: 1
Origin: null
Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7yyQ5XLHOn6WZ6MT
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
Connection: close

------WebKitFormBoundary7yyQ5XLHOn6WZ6MT
Content-Disposition: form-data; name="fname"; filename="#filename#"
Content-Type: image/jpeg

aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa?
------WebKitFormBoundary7yyQ5XLHOn6WZ6MT--

```

- 测试上传

```
uploadFuzz.exe -u http://127.0.0.1/saveupload.asp -i input.txt
```

