package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
)

func sayHellowName(w http.ResponseWriter, r *http.Request){
    r.ParseForm()   //解析参数，默认不会解析
    fmt.Println(r.Form)  //输出到服务端的打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("Scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])

    for k, v := range r.Form{
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

    fmt.Fprintf(w, "hello go web server")    //输出到客户端

}

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

func httpServer(port string){   
    // 监听绑定
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }    
}

func main() {
    fmt.Println("server start")
    // 注册函数，用户连接， 自动调用指定处理函数
    //http.HandleFunc("/hello", HelloServer)
    http.HandleFunc("/", sayHellowName)

    /*增加Http路径地址
    fsh := http.FileServer(http.Dir("/"))
    http.Handle("/", httpStripPrefix("/", fsh))
    */
    //go httpServer("9191")
    httpServer(":9090")
    
}

