# xiao-shorten-url
short url with bolt

### shorten
http://localhost:8000/api?long_url=https://www.baidu.com

```json
{
    "err_code": 0,
    "err_msg": "成功",
    "result": "http://localhost:8000/uYzU7r"
}
```

### redirect
http://localhost:8000/uYzU7r

```
http -f POST http://localhost:8000/api long_url=https://www.baidu.com

HTTP/1.1 200 OK
Content-Length: 74
Content-Type: application/json charset=utf-8
Date: Thu, 20 Jun 2019 03:55:40 GMT

{
    "err_code": 0,
    "err_msg": "成功",
    "result": "http://localhost:8000/uYzU7r"
}
```

```
http http://localhost:8000/uYzU7r

HTTP/1.1 303 See Other
Content-Length: 64
Content-Type: text/html; charset=utf-8
Date: Thu, 20 Jun 2019 03:56:06 GMT
Location: https://www.baidu.com

<a href="https://www.baidu.com">See Other</a>.
```
