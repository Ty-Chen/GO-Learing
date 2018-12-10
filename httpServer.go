package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

// hello world, the web server
// w: 给客户端回复数据， req: 读取客户端发送的数据
func HelloServer(w http.ResponseWriter, req *http.Request) {
    // 打印客户端头信息
    fmt.Println(req.Method)
    fmt.Println(req.Header)
    fmt.Println(req.Body)
    fmt.Println(req.URL)

    // 给客户端回复数据
    io.WriteString(w, "hello, world!\n")
    w.Write([]byte("lisa"))
}

func main() {
    // 注册函数，用户连接， 自动调用指定处理函数
    http.HandleFunc("/hello", HelloServer)

    // 监听绑定
    err := http.ListenAndServe("127.0.0.1:12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

